import type {
  DynamicUIComponent,
  InteractionRequest,
} from "@shopwise/protocols";
import { useState } from "react";

import type { AgentEnvelope } from "../api/decision-memory";

type DynamicUIRendererProps = Readonly<{
  envelope: AgentEnvelope;
  onConfirmCheckout?: (productId: string) => void;
  onInteraction?: (request: InteractionRequest) => void;
}>;

type Product = {
  id: string;
  name: string;
  price: number;
  match_explanation?: string;
  image?: string;
  specifications?: Record<string, unknown>;
};

const formatVND = (price: number) =>
  new Intl.NumberFormat("vi-VN", {
    style: "currency",
    currency: "VND",
  }).format(price);

const readProducts = (value: unknown): Product[] => {
  if (!Array.isArray(value)) {
    return [];
  }
  return value.filter(
    (product): product is Product =>
      typeof product === "object" &&
      product !== null &&
      typeof product.id === "string" &&
      typeof product.name === "string" &&
      typeof product.price === "number"
  );
};

const ProductList = ({ products }: Readonly<{ products: Product[] }>) => (
  <div className="grid gap-3 md:grid-cols-2 xl:grid-cols-3">
    {products.map((product) => (
      <article
        className="rounded-xl border border-white/15 bg-white/5 p-4"
        key={product.id}
      >
        {product.image ? (
          <img
            alt={product.name}
            className="mb-3 h-28 w-full rounded-lg object-cover"
            src={product.image}
          />
        ) : null}
        <h3 className="font-semibold text-white">{product.name}</h3>
        <p className="mt-1 font-mono text-sm text-emerald-400">
          {formatVND(product.price)}
        </p>
        {product.match_explanation ? (
          <p className="mt-2 text-sm text-white/70">
            {product.match_explanation}
          </p>
        ) : null}
      </article>
    ))}
  </div>
);

const ComponentRenderer = ({
  component,
}: Readonly<{ component: DynamicUIComponent }>) => {
  const children = component.children?.map((child) => (
    <ComponentRenderer component={child} key={child.id} />
  ));
  const title =
    typeof component.props?.title === "string" ? component.props.title : null;
  const message =
    typeof component.props?.message === "string"
      ? component.props.message
      : null;
  const products = readProducts(component.props?.products);

  switch (component.type) {
    case "page":
    case "section":
    case "grid":
    case "timeline":
      return <section className="grid gap-4">{children}</section>;
    case "card":
      return (
        <article className="rounded-xl border border-[#4F7CFF]/40 bg-[#11131a] p-5">
          {title ? <h3 className="font-semibold text-white">{title}</h3> : null}
          {children}
        </article>
      );
    case "product_card":
    case "product_carousel":
    case "product_comparison":
    case "checkout_summary":
      return (
        <section className="grid gap-4">
          {message ? <p className="text-white/80">{message}</p> : null}
          <ProductList products={products} />
          {children}
        </section>
      );
    case "chat_message":
    case "recommendation_explanation":
    case "decision_reasoning":
    case "promotion_banner":
    case "warranty_information":
    case "inventory_status":
    case "checkout_readiness":
      return (
        <p className="text-white/80">
          {message ?? title}
          {children}
        </p>
      );
    case "thinking_indicator":
      return <p aria-live="polite">Thinking…</p>;
    case "confidence_indicator":
      return (
        <meter max={100} min={0} value={Number(component.props?.value ?? 0)} />
      );
    case "specification_table":
      return <div className="grid gap-2">{children}</div>;
    case "tabs":
    case "modal":
    case "drawer":
      return <div className="grid gap-3">{children}</div>;
    default:
      // Unknown future components degrade to their known children.
      return children ? <>{children}</> : null;
  }
};

export const ClarificationCard = ({
  envelope,
  onInteraction,
}: Readonly<{
  envelope: Extract<AgentEnvelope, { type: "question" }>;
  onInteraction?: (request: InteractionRequest) => void;
}>) => {
  const [selectedOptionIds, setSelectedOptionIds] = useState<string[]>([]);
  const [freeText, setFreeText] = useState("");
  const question = envelope.question;

  if (!question) {
    return <p>{envelope.message}</p>;
  }

  const selectOption = (optionId: string) => {
    if (question.mode === "single") {
      setSelectedOptionIds([optionId]);
      return;
    }
    setSelectedOptionIds((current) =>
      current.includes(optionId)
        ? current.filter((id) => id !== optionId)
        : [...current, optionId]
    );
  };
  const submit = (event: React.SubmitEvent) => {
    event.preventDefault();
    if (selectedOptionIds.length === 0 && freeText.trim().length === 0) {
      return;
    }
    onInteraction?.({
      schema_version: "1.0",
      interaction_id: question.id,
      source_turn_id: envelope.turn_id ?? question.id,
      component_id: question.id,
      event: "submit",
      action: "question.answer",
      payload: {
        question_id: question.id,
        selected_option_ids: selectedOptionIds,
        free_text: freeText.trim(),
      },
    });
  };

  return (
    <form
      aria-label="Clarification question"
      className="mx-auto w-full max-w-xl rounded-2xl border border-[#4F7CFF]/40 bg-[#11131a] p-4 shadow-xl"
      onSubmit={submit}
    >
      <h2 className="text-md font-semibold text-white">{envelope.message}</h2>
      <fieldset className="mt-4 grid gap-2">
        <legend className="sr-only">Suggested answers</legend>
        {question.options.map((option) => (
          <label
            className="flex cursor-pointer items-center gap-3 rounded-lg border border-white/15 px-4 py-3 text-white/80 has-checked:border-[#4F7CFF] has-checked:bg-[#4F7CFF]/10"
            key={option.id}
          >
            <input
              checked={selectedOptionIds.includes(option.id)}
              name={question.id}
              onChange={() => selectOption(option.id)}
              type={question.mode === "single" ? "radio" : "checkbox"}
              value={option.id}
            />
            {option.label}
          </label>
        ))}
      </fieldset>
      {question.free_text_allowed ? (
        <label className="mt-4 grid gap-2 text-sm text-white/70">
          {question.input_label ?? "Your answer"}
          <input
            className="rounded-lg border border-white/20 bg-black/30 px-3 py-2 text-white outline-none focus:border-[#4F7CFF]"
            maxLength={1000}
            onChange={(event) => setFreeText(event.target.value)}
            placeholder={
              question.input_placeholder ?? "Add details in your own words"
            }
            type="text"
            value={freeText}
          />
        </label>
      ) : null}
      <button
        className="mt-5 rounded-lg bg-[#4F7CFF] px-5 py-2 font-semibold text-white disabled:cursor-not-allowed disabled:opacity-40"
        disabled={
          selectedOptionIds.length === 0 && freeText.trim().length === 0
        }
        type="submit"
      >
        {question.submit_label ?? "Submit"}
      </button>
    </form>
  );
};

const DynamicUIRenderer = ({
  envelope,
  onConfirmCheckout,
  onInteraction,
}: DynamicUIRendererProps) => {
  if (envelope.type === "question") {
    return (
      <ClarificationCard envelope={envelope} onInteraction={onInteraction} />
    );
  }

  if (envelope.ui_operations?.length) {
    return (
      <div className="grid gap-4">
        {envelope.ui_operations.map((operation) => (
          <ComponentRenderer
            component={operation.component as DynamicUIComponent}
            key={operation.id}
          />
        ))}
      </div>
    );
  }

  const isComparison = envelope.type === "comparison";
  const product = envelope.decision.products[0];
  return (
    <section
      aria-label={isComparison ? "Product comparison" : "Agent decision"}
    >
      <p>{envelope.message}</p>
      <ProductList products={envelope.decision.products} />
      {envelope.type === "checkout_ready" && product ? (
        <button type="button" onClick={() => onConfirmCheckout?.(product.id)}>
          Confirm checkout
        </button>
      ) : null}
    </section>
  );
};

export default DynamicUIRenderer;

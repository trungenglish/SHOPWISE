import type { AgentEnvelope } from "../api/decision-memory";

type DynamicUIRendererProps = Readonly<{
  envelope: AgentEnvelope;
  onConfirmCheckout?: (productId: string) => void;
}>;

const formatVND = (price: number) =>
  new Intl.NumberFormat("vi-VN", {
    style: "currency",
    currency: "VND",
  }).format(price);

const DynamicUIRenderer = ({
  envelope,
  onConfirmCheckout,
}: DynamicUIRendererProps) => {
  if (envelope.type === "question") {
    return <p>{envelope.message}</p>;
  }

  const isComparison = envelope.type === "comparison";
  const product = envelope.decision.products[0];

  return (
    <section
      aria-label={isComparison ? "Product comparison" : "Agent decision"}
    >
      <p>{envelope.message}</p>
      {isComparison ? (
        <table>
          <tbody>
            {envelope.decision.products.map((item) => (
              <tr key={item.id}>
                <th scope="row">{item.name}</th>
                <td>{formatVND(item.price)}</td>
                <td>{item.match_explanation}</td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        <ul>
          {envelope.decision.products.map((item) => (
            <li key={item.id}>
              <h3>{item.name}</h3>
              <p>{formatVND(item.price)}</p>
              <p>{item.match_explanation}</p>
            </li>
          ))}
        </ul>
      )}
      {envelope.type === "checkout_ready" && product ? (
        <button type="button" onClick={() => onConfirmCheckout?.(product.id)}>
          Confirm checkout
        </button>
      ) : null}
    </section>
  );
};

export default DynamicUIRenderer;

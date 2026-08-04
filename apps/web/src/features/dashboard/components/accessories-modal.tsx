import { useQuery } from "@tanstack/react-query";
import { Check, ExternalLink, X } from "lucide-react";
import { useMemo, useRef, useState } from "react";

import {
  type AccessoryRecommendation,
  getAccessoryRecommendations,
} from "@/api/accessories";

import type { Laptop } from "../types";

interface AccessoriesModalProps {
  isOpen: boolean;
  onClose: () => void;
  products: Laptop[];
  onAddToCheckout: (accessory: AccessoryRecommendation) => void;
}

const ALL = "all";
const PAGE_SIZE = 12;

const formatVND = (amount: number) =>
  new Intl.NumberFormat("vi-VN", {
    style: "currency",
    currency: "VND",
  }).format(amount);

export default function AccessoriesModal({
  isOpen,
  onClose,
  products,
  onAddToCheckout,
}: AccessoriesModalProps) {
  const [productFilter, setProductFilter] = useState(ALL);
  const [categoryFilter, setCategoryFilter] = useState(ALL);
  const [kindFilter, setKindFilter] = useState(ALL);
  const [currentPage, setCurrentPage] = useState(1);
  const listRef = useRef<HTMLDivElement>(null);
  const productIds = useMemo(
    () => products.map((product) => product.id),
    [products]
  );
  const recommendations = useQuery({
    queryKey: ["accessory-recommendations", productIds],
    queryFn: () => getAccessoryRecommendations(productIds),
    enabled: isOpen && productIds.length > 0,
  });
  const items = recommendations.data ?? [];
  const categories = useMemo(
    () => Array.from(new Set(items.map((item) => item.category))).sort(),
    [items]
  );
  const filteredItems = useMemo(
    () =>
      items.filter(
        (item) =>
          (productFilter === ALL ||
            item.compatible_product_ids.includes(productFilter)) &&
          (categoryFilter === ALL || item.category === categoryFilter) &&
          (kindFilter === ALL || item.kind === kindFilter)
      ),
    [categoryFilter, items, kindFilter, productFilter]
  );
  const productNames = useMemo(
    () => new Map(products.map((product) => [product.id, product.name])),
    [products]
  );
  const pageCount = Math.max(1, Math.ceil(filteredItems.length / PAGE_SIZE));
  const visibleItems = filteredItems.slice(
    (currentPage - 1) * PAGE_SIZE,
    currentPage * PAGE_SIZE
  );
  const changePage = (page: number) => {
    setCurrentPage(page);
    if (listRef.current) {
      listRef.current.scrollTop = 0;
    }
  };

  if (!isOpen) {
    return null;
  }

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/75 p-4 backdrop-blur-md">
      <div className="bg-surface-high border-outline-variant/30 flex max-h-[85vh] w-full max-w-4xl flex-col overflow-hidden rounded-2xl border shadow-2xl">
        <header className="border-outline-variant/15 bg-surface-lowest border-b p-5">
          <div className="flex items-start justify-between gap-4">
            <div>
              <h3 className="text-on-surface font-display text-lg font-bold">
                Phụ kiện đề xuất từ Phong Vũ
              </h3>
              <p className="text-on-surface-variant mt-1 text-xs">
                Giá và tồn kho được đồng bộ từ nguồn Phong Vũ.
              </p>
            </div>
            <button
              type="button"
              onClick={onClose}
              aria-label="Đóng modal phụ kiện"
              className="text-on-surface-variant hover:text-on-surface rounded-lg p-1.5"
            >
              <X size={18} />
            </button>
          </div>
          <div className="mt-4 grid gap-3 sm:grid-cols-3">
            <label className="text-on-surface-variant text-xs">
              Lọc theo laptop
              <select
                value={productFilter}
                onChange={(event) => {
                  setProductFilter(event.target.value);
                  changePage(1);
                }}
                className="bg-surface border-outline-variant/30 text-on-surface mt-1 w-full rounded-lg border p-2"
              >
                <option value={ALL}>Tất cả laptop đề xuất</option>
                {products.map((product) => (
                  <option key={product.id} value={product.id}>
                    {product.name}
                  </option>
                ))}
              </select>
            </label>
            <label className="text-on-surface-variant text-xs">
              Loại phụ kiện
              <select
                value={kindFilter}
                onChange={(event) => {
                  setKindFilter(event.target.value);
                  changePage(1);
                }}
                className="bg-surface border-outline-variant/30 text-on-surface mt-1 w-full rounded-lg border p-2"
              >
                <option value={ALL}>Tất cả</option>
                <option value="external">Phụ kiện ngoài</option>
                <option value="upgrade">Nâng cấp</option>
              </select>
            </label>
            <label className="text-on-surface-variant text-xs">
              Danh mục
              <select
                value={categoryFilter}
                onChange={(event) => {
                  setCategoryFilter(event.target.value);
                  changePage(1);
                }}
                className="bg-surface border-outline-variant/30 text-on-surface mt-1 w-full rounded-lg border p-2"
              >
                <option value={ALL}>Tất cả</option>
                {categories.map((category) => (
                  <option key={category} value={category}>
                    {category}
                  </option>
                ))}
              </select>
            </label>
          </div>
        </header>

        <div
          ref={listRef}
          className="grid min-h-0 flex-1 content-start gap-4 overflow-auto p-5 md:grid-cols-2"
        >
          {recommendations.isLoading && (
            <p className="text-on-surface-variant col-span-full py-10 text-center">
              Đang tải dữ liệu Phong Vũ…
            </p>
          )}
          {recommendations.isError && (
            <div className="col-span-full py-10 text-center">
              <p className="text-red-400">
                Không thể tải phụ kiện từ Phong Vũ.
              </p>
              <button
                type="button"
                onClick={() => recommendations.refetch()}
                className="text-primary mt-3 underline"
              >
                Thử lại
              </button>
            </div>
          )}
          {!recommendations.isLoading &&
            !recommendations.isError &&
            filteredItems.length === 0 && (
              <p className="text-on-surface-variant col-span-full py-10 text-center">
                Không có phụ kiện phù hợp với bộ lọc này.
              </p>
            )}
          {visibleItems.map((item) => {
            const applicableModels = item.compatible_product_ids
              .map((id) => productNames.get(id))
              .filter((name): name is string => Boolean(name));
            return (
              <article
                key={item.id}
                className="bg-surface-low border-outline-variant/20 flex h-max overflow-hidden rounded-xl border"
              >
                <div className="bg-primary/5 flex w-28 shrink-0 items-center justify-center overflow-hidden">
                  {item.image_url ? (
                    <img
                      src={item.image_url}
                      alt={item.name}
                      className="h-full w-full object-cover"
                    />
                  ) : (
                    <Check className="text-primary" size={24} />
                  )}
                </div>
                <div className="min-w-0 flex-1 p-4">
                  <div className="flex items-start justify-between gap-3">
                    <div>
                      <span className="text-primary font-mono text-[10px] uppercase">
                        {item.kind === "upgrade"
                          ? "Nâng cấp"
                          : "Phụ kiện ngoài"}{" "}
                        · {item.category}
                      </span>
                      <h4 className="text-on-surface mt-1 text-sm font-bold">
                        {item.name}
                      </h4>
                    </div>
                    <strong className="text-on-surface shrink-0 text-sm">
                      {formatVND(item.price)}
                    </strong>
                  </div>
                  <p className="text-on-surface-variant mt-2 text-xs">
                    {item.compatibility_reason}
                  </p>
                  <p className="text-on-surface-variant mt-1 text-[11px]">
                    Áp dụng: {applicableModels.join(", ")}
                  </p>
                  <p className="text-on-surface-variant mt-1 text-[11px]">
                    {item.in_stock ? "Còn hàng" : "Hết hàng"} · cập nhật{" "}
                    {new Date(item.fetched_at).toLocaleString("vi-VN")}
                  </p>
                  {item.offer_state === "stale" && (
                    <p className="mt-1 text-[11px] text-amber-400">
                      Dữ liệu giá đã cũ
                    </p>
                  )}
                  <div className="mt-3 flex flex-wrap gap-2">
                    <a
                      href={item.source_url}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="border-outline-variant/30 text-on-surface inline-flex items-center gap-1 rounded-md border px-2.5 py-1.5 text-xs"
                    >
                      Xem tại Phong Vũ <ExternalLink size={12} />
                    </a>
                    <button
                      type="button"
                      disabled={!item.checkout_available}
                      onClick={() => onAddToCheckout(item)}
                      className="bg-primary text-on-primary rounded-md px-2.5 py-1.5 text-xs disabled:cursor-not-allowed disabled:opacity-40"
                    >
                      Thêm vào checkout
                    </button>
                  </div>
                </div>
              </article>
            );
          })}
        </div>
        {!recommendations.isLoading &&
          !recommendations.isError &&
          filteredItems.length > 0 && (
            <nav
              aria-label="Phân trang phụ kiện"
              className="border-outline-variant/15 flex items-center justify-center gap-4 border-t p-4"
            >
              <button
                type="button"
                disabled={currentPage === 1}
                onClick={() => changePage(currentPage - 1)}
                className="border-outline-variant/30 text-on-surface rounded-md border px-3 py-1.5 text-xs disabled:cursor-not-allowed disabled:opacity-40"
              >
                Trước
              </button>
              <span
                className="text-on-surface-variant text-xs"
                aria-live="polite"
              >
                Trang {currentPage}/{pageCount}
              </span>
              <button
                type="button"
                disabled={currentPage === pageCount}
                onClick={() => changePage(currentPage + 1)}
                className="border-outline-variant/30 text-on-surface rounded-md border px-3 py-1.5 text-xs disabled:cursor-not-allowed disabled:opacity-40"
              >
                Sau
              </button>
            </nav>
          )}
      </div>
    </div>
  );
}

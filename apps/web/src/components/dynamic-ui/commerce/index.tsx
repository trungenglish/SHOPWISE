import React from 'react';
import type {
  ProductCardProps,
  ProductCarouselProps,
  ProductComparisonProps,
  SpecificationTableProps,
  PromotionBannerProps,
  WarrantyInformationProps,
  InventoryStatusProps,
  CheckoutSummaryProps
} from '@shopwise/ui-protocol/types/commerce';

const formatVND = (amount: number) => {
  return new Intl.NumberFormat('vi-VN', { style: 'currency', currency: 'VND' }).format(amount);
};

export const ProductCard: React.FC<ProductCardProps> = ({ name, priceVND, imageUrl, rating }) => (
  <div className="product-card">
    {imageUrl && <img src={imageUrl} alt={name} className="product-image" />}
    <h3 className="product-name">{name}</h3>
    <p className="product-price">{formatVND(priceVND)}</p>
    {rating !== undefined && <div className="product-rating">Rating: {rating}/5</div>}
  </div>
);

export const ProductCarousel: React.FC<ProductCarouselProps & { children?: React.ReactNode }> = ({ title, children }) => (
  <div className="product-carousel">
    {title && <h2>{title}</h2>}
    <div className="carousel-items">
      {children}
    </div>
  </div>
);

export const ProductComparison: React.FC<ProductComparisonProps> = ({ features }) => (
  <div className="product-comparison">
    {/* MVP Placeholder */}
    <p>Comparing products by: {features?.join(', ')}</p>
  </div>
);

export const SpecificationTable: React.FC<SpecificationTableProps> = ({ specs }) => (
  <table className="specification-table">
    <tbody>
      {specs && Object.entries(specs).map(([key, value]) => (
        <tr key={key}>
          <th>{key}</th>
          <td>{value}</td>
        </tr>
      ))}
    </tbody>
  </table>
);

export const PromotionBanner: React.FC<PromotionBannerProps> = ({ text, discountCode }) => (
  <div className="promotion-banner">
    {text && <p>{text}</p>}
    {discountCode && <strong>Code: {discountCode}</strong>}
  </div>
);

export const WarrantyInformation: React.FC<WarrantyInformationProps> = ({ months, details }) => (
  <div className="warranty-info">
    {months !== undefined && <p>Warranty: {months} months</p>}
    {details && <p>{details}</p>}
  </div>
);

export const InventoryStatus: React.FC<InventoryStatusProps> = ({ inStock, quantity }) => (
  <div className="inventory-status">
    {inStock ? (
      <span className="in-stock">In Stock {quantity !== undefined ? `(${quantity} left)` : ''}</span>
    ) : (
      <span className="out-of-stock">Out of Stock</span>
    )}
  </div>
);

export const CheckoutSummary: React.FC<CheckoutSummaryProps> = ({ items, totalPriceVND }) => (
  <div className="checkout-summary">
    <h3>Checkout Summary</h3>
    <ul>
      {items?.map((item, idx) => (
        <li key={idx}>
          {item.productId} x {item.quantity} = {item.priceVND !== undefined ? formatVND(item.priceVND) : ''}
        </li>
      ))}
    </ul>
    {totalPriceVND !== undefined && (
      <div className="total">
        <strong>Total: {formatVND(totalPriceVND)}</strong>
      </div>
    )}
  </div>
);

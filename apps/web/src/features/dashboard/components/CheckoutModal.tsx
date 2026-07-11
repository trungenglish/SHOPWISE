import { useState } from "react";
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import {
  X,
  ShoppingBag,
  CreditCard,
  Sparkles,
  Check,
  Truck,
  Store,
  Loader2,
  Search,
  MapPin
} from "lucide-react";
import { Laptop } from "../types";

import { Button } from "@shopwise/ui/components/button";
import { Input } from "@shopwise/ui/components/input";
import {
  Field,
  FieldLabel,
  FieldError,
  FieldContent,
} from "@shopwise/ui/components/field";
import {
  RadioGroup,
  RadioGroupItem,
} from "@shopwise/ui/components/radio-group";
import { Label } from "@shopwise/ui/components/label";

interface CheckoutModalProps {
  isOpen: boolean;
  onClose: () => void;
  product: Laptop | null;
  discountRate: number; // calculated from connected retail accounts
}

const checkoutSchema = z
  .object({
    fullName: z.string().min(1, "Vui lòng nhập họ và tên"),
    phoneNumber: z
      .string()
      .trim()
      .min(1, "Vui lòng nhập số điện thoại")
      .regex(
        /^(03|05|07|08|09)\d{8}$/,
        "Số điện thoại không hợp lệ (vd: 0912345678)"
      ),
    email: z
      .string()
      .min(1, "Vui lòng nhập email")
      .email("Email không hợp lệ"),
    deliveryMethod: z.enum(["delivery", "pickup"]),
    address: z.string().optional(),
    storeId: z.string().optional(),
  })
  .superRefine((data, ctx) => {
    if (
      data.deliveryMethod === "delivery" &&
      (!data.address || data.address.trim() === "")
    ) {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        message: "Vui lòng nhập địa chỉ giao hàng",
        path: ["address"],
      });
    }
    if (
      data.deliveryMethod === "pickup" &&
      (!data.storeId || data.storeId.trim() === "")
    ) {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        message: "Vui lòng chọn một cửa hàng",
        path: ["storeId"],
      });
    }
  });

type CheckoutFormData = z.infer<typeof checkoutSchema>;

const formatVND = (amount: number) => {
  const vndAmount = amount;
  return new Intl.NumberFormat("vi-VN", {
    style: "currency",
    currency: "VND",
  }).format(vndAmount);
};

const MOCK_STORES = [
  { id: "pv-1", name: "Phong Vũ Nguyễn Thị Minh Khai", address: "264 Nguyễn Thị Minh Khai, P. Võ Thị Sáu, Q. 3, TP.HCM" },
  { id: "pv-2", name: "Phong Vũ Cách Mạng Tháng 8", address: "288 Cách Mạng Tháng 8, P. 10, Q. 3, TP.HCM" },
  { id: "pv-3", name: "Phong Vũ Thái Hà", address: "1 Thái Hà, Trung Liệt, Đống Đa, Hà Nội" },
  { id: "pv-4", name: "Phong Vũ Cầu Giấy", address: "173 Xuân Thủy, Dịch Vọng Hậu, Cầu Giấy, Hà Nội" },
  { id: "pv-5", name: "Phong Vũ Lê Văn Việt", address: "1A Lê Văn Việt, Hiệp Phú, TP. Thủ Đức, TP.HCM" },
];

export default function CheckoutModal({
  isOpen,
  onClose,
  product,
  discountRate,
}: CheckoutModalProps) {
  const [promoCode, setPromoCode] = useState("");
  const [promoApplied, setPromoApplied] = useState(false);
  const [paymentSuccess, setPaymentSuccess] = useState(false);
  const [isSubmittingManual, setIsSubmittingManual] = useState(false);
  const [storeSearch, setStoreSearch] = useState("");

  const filteredStores = MOCK_STORES.filter(s => 
    s.name.toLowerCase().includes(storeSearch.toLowerCase()) || 
    s.address.toLowerCase().includes(storeSearch.toLowerCase())
  );

  const {
    register,
    handleSubmit,
    watch,
    control,
    formState: { errors },
  } = useForm<CheckoutFormData>({
    resolver: zodResolver(checkoutSchema),
    defaultValues: {
      fullName: "",
      phoneNumber: "",
      email: "",
      deliveryMethod: "delivery",
      address: "",
      storeId: "",
    },
    mode: "onTouched",
  });

  const deliveryMethod = watch("deliveryMethod");

  if (!isOpen) return null;
  if (!product) return null;

  const basePrice = product.price;
  const retailDiscount = basePrice * discountRate;
  const promoDiscount = promoApplied ? basePrice * 0.05 : 0; // extra 5% for promo "SHOPWISE5"
  const shipping = basePrice > 37500000 ? 0 : 625000;
  const tax = (basePrice - retailDiscount - promoDiscount) * 0.08;
  const finalTotal =
    basePrice - retailDiscount - promoDiscount + shipping + tax;

  const handleApplyPromo = () => {
    if (promoCode.toUpperCase() === "SHOPWISE5") {
      setPromoApplied(true);
      setPromoCode("");
    } else {
      alert("Mã giảm giá không hợp lệ. Hãy thử: SHOPWISE5");
    }
  };

  const handlePayment = async (data: CheckoutFormData) => {
    setIsSubmittingManual(true);
    // Simulate network delay
    await new Promise((resolve) => setTimeout(resolve, 1500));
    setPaymentSuccess(true);
    setIsSubmittingManual(false);
  };

  return (
    <div className="bg-black/80 fixed inset-0 z-50 flex items-center justify-center p-4 select-none backdrop-blur-sm">
      <div
        className="bg-surface-high border-outline-variant/30 flex w-full max-w-4xl flex-col overflow-hidden rounded-xl border shadow-2xl"
        style={{ maxHeight: "calc(100vh - 2rem)" }}
      >
        {/* Header */}
        <div className="border-outline-variant/20 bg-surface flex items-center justify-between border-b p-4 md:px-6">
          <div className="flex items-center gap-3">
            <ShoppingBag size={20} className="text-primary" />
            <div>
              <h2 className="text-on-surface text-lg font-bold">
                Hoàn tất đơn hàng
              </h2>
              <p className="text-on-surface-variant text-sm">
                Kiểm tra thông tin giao hàng và xác nhận đơn hàng của bạn.
              </p>
            </div>
          </div>
          <button
            onClick={onClose}
            className="text-on-surface-variant hover:text-on-surface hover:bg-surface-highest cursor-pointer rounded-lg p-2 transition-colors"
            aria-label="Close modal"
          >
            <X size={20} />
          </button>
        </div>

        {/* Content */}
        <div className="flex-1 overflow-y-auto">
          {paymentSuccess ? (
            <div className="flex flex-col items-center gap-4 py-16 text-center">
              <div className="flex h-16 w-16 items-center justify-center rounded-full border border-green-500/30 bg-green-500/10 text-green-500">
                <Check size={32} />
              </div>
              <div>
                <h4 className="text-on-surface font-display text-xl font-bold">
                  Đặt hàng thành công!
                </h4>
                <p className="text-on-surface-variant mx-auto mt-2 max-w-sm text-sm">
                  Đơn hàng <strong>{product.name}</strong> của bạn đã được xác
                  nhận. Nhân viên sẽ liên hệ với bạn trong vòng 15 phút.
                </p>
              </div>
              <Button
                onClick={() => {
                  setPaymentSuccess(false);
                  onClose();
                }}
                className="mt-6 px-8"
              >
                Đóng
              </Button>
            </div>
          ) : (
            <form
              onSubmit={handleSubmit(handlePayment)}
              className="flex h-full flex-col md:grid md:grid-cols-[1fr_380px]"
            >
              {/* Left Column */}
              <div className="border-outline-variant/20 flex flex-col gap-8 p-5 md:border-r md:p-8">
                {/* 1. Thông tin khách hàng */}
                <section>
                  <h3 className="text-primary mb-5 text-sm font-bold tracking-wide uppercase">
                    1. Thông tin khách hàng
                  </h3>
                  <div className="flex flex-col gap-4">
                    <Field data-invalid={!!errors.fullName}>
                      <FieldLabel htmlFor="fullName">Họ và tên</FieldLabel>
                      <FieldContent>
                        <Input
                          id="fullName"
                          placeholder="Nguyễn Văn A"
                          autoComplete="name"
                          disabled={isSubmittingManual}
                          {...register("fullName")}
                          aria-invalid={!!errors.fullName}
                        />
                        <FieldError errors={[errors.fullName]} />
                      </FieldContent>
                    </Field>
                    <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
                      <Field data-invalid={!!errors.phoneNumber}>
                        <FieldLabel htmlFor="phoneNumber">
                          Số điện thoại
                        </FieldLabel>
                        <FieldContent>
                          <Input
                            id="phoneNumber"
                            type="tel"
                            placeholder="0912345678"
                            autoComplete="tel"
                            disabled={isSubmittingManual}
                            {...register("phoneNumber")}
                            aria-invalid={!!errors.phoneNumber}
                          />
                          <FieldError errors={[errors.phoneNumber]} />
                        </FieldContent>
                      </Field>
                      <Field data-invalid={!!errors.email}>
                        <FieldLabel htmlFor="email">Email</FieldLabel>
                        <FieldContent>
                          <Input
                            id="email"
                            type="email"
                            placeholder="you@example.com"
                            autoComplete="email"
                            disabled={isSubmittingManual}
                            {...register("email")}
                            aria-invalid={!!errors.email}
                          />
                          <FieldError errors={[errors.email]} />
                        </FieldContent>
                      </Field>
                    </div>
                  </div>
                </section>

                {/* 2. Hình thức nhận hàng */}
                <section>
                  <h3 className="text-primary mb-5 text-sm font-bold tracking-wide uppercase">
                    2. Hình thức nhận hàng
                  </h3>
                  <Controller
                    control={control}
                    name="deliveryMethod"
                    render={({ field }) => (
                      <RadioGroup
                        value={field.value}
                        onValueChange={field.onChange}
                        className="mb-5 grid grid-cols-1 gap-3 sm:grid-cols-2"
                        disabled={isSubmittingManual}
                      >
                        <Label
                          htmlFor="delivery"
                          className="border-outline-variant/30 hover:border-primary/50 has-[[data-state=checked]]:border-primary flex cursor-pointer items-center space-x-3 rounded-lg border p-4 transition-colors"
                        >
                          <RadioGroupItem value="delivery" id="delivery" />
                          <span className="flex items-center gap-2 text-sm">
                            <Truck size={16} /> Giao tận nơi
                          </span>
                        </Label>
                        <Label
                          htmlFor="pickup"
                          className="border-outline-variant/30 hover:border-primary/50 has-[[data-state=checked]]:border-primary flex cursor-pointer items-center space-x-3 rounded-lg border p-4 transition-colors"
                        >
                          <RadioGroupItem value="pickup" id="pickup" />
                          <span className="flex items-center gap-2 text-sm">
                            <Store size={16} /> Đến lấy tại store
                          </span>
                        </Label>
                      </RadioGroup>
                    )}
                  />

                  {deliveryMethod === "delivery" && (
                    <div className="animate-in fade-in slide-in-from-top-2">
                      <Field data-invalid={!!errors.address}>
                        <FieldLabel htmlFor="address">
                          Địa chỉ giao hàng
                        </FieldLabel>
                        <FieldContent>
                          <Input
                            id="address"
                            placeholder="Số nhà, tên đường, phường/xã, quận/huyện, tỉnh/thành phố"
                            autoComplete="street-address"
                            disabled={isSubmittingManual}
                            {...register("address")}
                            aria-invalid={!!errors.address}
                          />
                          <FieldError errors={[errors.address]} />
                        </FieldContent>
                      </Field>
                    </div>
                  )}
                  {deliveryMethod === "pickup" && (
                    <div className="animate-in fade-in slide-in-from-top-2 flex flex-col gap-4">
                      <div className="relative">
                        <Search className="text-on-surface-variant absolute top-1/2 left-3 -translate-y-1/2" size={16} />
                        <Input
                          placeholder="Tìm kiếm cửa hàng theo tên hoặc địa chỉ..."
                          value={storeSearch}
                          onChange={(e) => setStoreSearch(e.target.value)}
                          className="pl-9"
                          disabled={isSubmittingManual}
                        />
                      </div>
                      
                      <div className="border-outline-variant/30 flex max-h-[250px] flex-col overflow-y-auto overflow-x-hidden rounded-lg border">
                        <Controller
                          control={control}
                          name="storeId"
                          render={({ field }) => (
                            <RadioGroup
                              value={field.value}
                              onValueChange={field.onChange}
                              className="flex flex-col gap-0"
                              disabled={isSubmittingManual}
                            >
                              {filteredStores.length > 0 ? filteredStores.map((store) => (
                                <Label
                                  key={store.id}
                                  htmlFor={store.id}
                                  className="border-outline-variant/20 hover:bg-surface-low has-[[data-state=checked]]:bg-primary/5 flex cursor-pointer items-start gap-3 border-b p-4 transition-colors last:border-0"
                                >
                                  <RadioGroupItem value={store.id} id={store.id} className="mt-0.5" />
                                  <div className="flex flex-col gap-1">
                                    <span className="text-on-surface text-sm font-bold">{store.name}</span>
                                    <span className="text-on-surface-variant flex items-start gap-1 text-xs">
                                      <MapPin size={14} className="mt-0.5 shrink-0" />
                                      {store.address}
                                    </span>
                                  </div>
                                </Label>
                              )) : (
                                <div className="text-on-surface-variant p-4 text-center text-sm">
                                  Không tìm thấy cửa hàng nào phù hợp.
                                </div>
                              )}
                            </RadioGroup>
                          )}
                        />
                      </div>
                      <FieldError errors={[errors.storeId]} />
                    </div>
                  )}
                </section>
              </div>

              {/* Right Column */}
              <div className="bg-surface-lowest flex flex-col p-5 md:p-8">
                <div className="bg-surface-high border-outline-variant/30 mb-8 flex items-center gap-4 rounded-xl border p-4 shadow-sm">
                  <img
                    src={product.image}
                    alt={product.name}
                    className="border-outline-variant/20 h-16 w-24 rounded-lg border object-cover"
                  />
                  <div className="min-w-0 flex-1">
                    <h4 className="text-on-surface truncate text-sm font-bold">
                      {product.name}
                    </h4>
                    <p className="text-primary mt-1 truncate text-xs font-medium">
                      {product.specs.gpu}
                    </p>
                  </div>
                  <span className="text-on-surface shrink-0 text-sm font-bold">
                    {formatVND(basePrice)}
                  </span>
                </div>

                <div className="mb-8 flex gap-2">
                  <Input
                    value={promoCode}
                    onChange={(e) => setPromoCode(e.target.value)}
                    placeholder="Mã giảm giá (SHOPWISE5)..."
                    disabled={isSubmittingManual}
                    className="flex-1"
                  />
                  <Button
                    type="button"
                    variant="secondary"
                    onClick={handleApplyPromo}
                    disabled={isSubmittingManual}
                  >
                    Áp dụng
                  </Button>
                </div>

                <div className="mb-8 flex flex-col gap-4">
                  <div className="text-on-surface-variant flex justify-between text-sm">
                    <span>Giá gốc sản phẩm</span>
                    <span className="text-on-surface font-medium">
                      {formatVND(basePrice)}
                    </span>
                  </div>

                  <div className="text-on-surface-variant flex justify-between text-sm">
                    <span>Phí giao hàng bảo đảm chống sốc</span>
                    <span className="text-on-surface font-medium">
                      {shipping === 0 ? "MIỄN PHÍ" : formatVND(shipping)}
                    </span>
                  </div>

                  {retailDiscount > 0 && (
                    <div className="flex justify-between text-sm text-green-500">
                      <span>Giảm giá thành viên</span>
                      <span className="font-medium">
                        -{formatVND(retailDiscount)}
                      </span>
                    </div>
                  )}

                  {promoApplied && (
                    <div className="flex justify-between text-sm text-green-500">
                      <span>Mã giảm giá</span>
                      <span className="font-medium">
                        -{formatVND(promoDiscount)}
                      </span>
                    </div>
                  )}

                  {tax > 0 && (
                    <div className="text-on-surface-variant flex justify-between text-sm">
                      <span>Thuế đối soát hải quan (8%)</span>
                      <span className="text-on-surface font-medium">
                        {formatVND(tax)}
                      </span>
                    </div>
                  )}

                  <div className="border-outline-variant/30 mt-2 flex justify-between border-t pt-5">
                    <span className="flex items-center gap-1.5 text-base font-bold text-primary">
                      <Sparkles size={16} /> Tổng thanh toán
                    </span>
                    <span className="text-lg font-bold text-primary">
                      {formatVND(finalTotal)}
                    </span>
                  </div>
                </div>

                <div className="mt-auto pt-4">
                  <Button
                    type="submit"
                    size="lg"
                    className="w-full text-base font-bold"
                    disabled={isSubmittingManual}
                  >
                    {isSubmittingManual ? (
                      <span
                        className="flex items-center gap-2"
                        aria-live="polite"
                      >
                        <Loader2 className="h-5 w-5 animate-spin" />
                        Đang xử lý...
                      </span>
                    ) : (
                      <span className="flex items-center gap-2">
                        <CreditCard size={18} /> Xác nhận đặt hàng
                      </span>
                    )}
                  </Button>
                </div>
              </div>
            </form>
          )}
        </div>
      </div>
    </div>
  );
}

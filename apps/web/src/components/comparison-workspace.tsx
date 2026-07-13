export interface ComparisonWorkspaceProps {
  data: any;
  onReplaceProduct?: (oldId: string, newId: string) => void;
  onRemoveProduct?: (id: string) => void;
  onProceedToCheckout?: (id: string) => void;
}

export function ComparisonWorkspace({ data, onReplaceProduct, onRemoveProduct, onProceedToCheckout }: ComparisonWorkspaceProps) {
  return (
    <div className="comparison-workspace p-4 border rounded shadow-sm">
      <h2 className="text-xl font-bold mb-4">Product Comparison</h2>
      <div className="flex gap-4 mb-4">
        {data?.products?.map((product: any) => {
           const isOutOfStock = product.availability === "OUT_OF_STOCK";
           return (
             <div key={product.id} className={`border p-4 rounded bg-card text-card-foreground ${isOutOfStock ? 'opacity-50' : ''}`}>
               <h3 className="font-bold text-lg">{product.name || product.id}</h3>
               {isOutOfStock && <p className="text-destructive text-sm">Out of Stock</p>}
               <button 
                  onClick={() => onProceedToCheckout?.(product.id)}
                  disabled={isOutOfStock}
                  className="mt-4 bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2 rounded-md font-medium w-full disabled:opacity-50 disabled:cursor-not-allowed"
               >
                  Proceed to Checkout
               </button>
             </div>
           );
        })}
      </div>
      <pre className="bg-muted p-4 rounded overflow-auto">{JSON.stringify(data, null, 2)}</pre>
    </div>
  );
}

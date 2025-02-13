package source

const (
	PRODUCT_INSERT       = "INSERT INTO products(name, sku) VALUES(?,?)"
	PRODUCT_COUNT_BY_SKU = "SELECT COUNT(*) FROM products WHERE sku = ?"
	PRODUCT_FIND_BY_ID   = "SELECT id, name, sku, stock, created_at, updated_at FROM products WHERE id = ?"
	PRODUCT_DELETE       = "DELETE products WHERE id = ?"
)

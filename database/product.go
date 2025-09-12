package database

// product struct
type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "orange is orange. I like orange",
		Price:       100,
		ImgURL:      "https://imgs.search.brave.com/0L7DFeQB0UMo_ytgjJ4EPf87gEccQDL6uiyMVzuTI34/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9zdGF0/aWMudmVjdGVlenku/Y29tL3N5c3RlbS9y/ZXNvdXJjZXMvdGh1/bWJuYWlscy8wMTgv/OTI3LzE2MS9zbWFs/bC9vcmFuZ2UtZnJ1/aXQtY29sb3ItaWxs/dXN0cmF0aW9uLXBu/Zy5wbmc",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "orange is orange. I like orange",
		Price:       100,
		ImgURL:      "https://imgs.search.brave.com/5_BRcLV8qhdtGjsjGHYkvaWpA0GXvd6V5L7-AL-zxZc/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9zdGF0/aWMudmVjdGVlenku/Y29tL3N5c3RlbS9y/ZXNvdXJjZXMvdGh1/bWJuYWlscy8wMjQv/MjE1LzExNC9zbWFs/bC9yaXBlLWFwcGxl/cy1mcnVpdC13aXRo/LWxlYXZlcy1waG90/by5qcGc",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "orange is orange. I like orange",
		Price:       100,
		ImgURL:      "https://imgs.search.brave.com/ZIwSNRtZ5AVAniDBDU3EpGJLMQc_1G-neNobwooYqYE/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly90aHVt/YnMuZHJlYW1zdGlt/ZS5jb20vYi9idW5j/aC1iYW5hbmEtZnJ1/aXRzLWN1dC1iYW5h/bmFzLWlzb2xhdGVk/LXdoaXRlLWJhY2tn/cm91bmQtYnVuY2gt/YmFuYW5hLWZydWl0/cy1jdXQtYmFuYW5h/cy1pc29sYXRlZC0x/MDM2NzUzMTkuanBn",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Grapes",
		Description: "orange is orange. I like orange",
		Price:       100,
		ImgURL:      "https://imgs.search.brave.com/agNMnjQSLtnTyMIddcz4GiaDYlx2T7G7ORT4EMtu9z4/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9zdGF0/aWMudmVjdGVlenku/Y29tL3N5c3RlbS9y/ZXNvdXJjZXMvdGh1/bWJuYWlscy8wMDYv/OTgyLzE1NC9zbWFs/bC9mcmVzaC1yZWQt/b3ItcHVycGxlLWdy/YXBlcy1mcnVpdC1p/c29sYXRlZC1vbi13/aGl0ZS1iYWNrZ3Jv/dW5kLXBob3RvLmpw/Zw",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Mango",
		Description: "orange is orange. I like orange",
		Price:       100,
		ImgURL:      "https://imgs.search.brave.com/PBhQraLJWOy5Aerb6kNwGoRf-uNCFyaUb38CMifp2O0/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly90aHVt/YnMuZHJlYW1zdGlt/ZS5jb20vYi9tYW5n/by1mcnVpdC1tYW5n/by1jdWJlcy13b29k/ZW4tdGFibGUtNjA1/OTc4ODcuanBn",
	}
	prd6 := Product{
		ID:          6,
		Title:       "Lichi",
		Description: "orange is orange. I like orange",
		Price:       100,
		ImgURL:      "https://imgs.search.brave.com/wYwhLuzjJLRovZpxM3gS9mTbale-u7jBx_lN7Hlq2F0/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9tZWRp/YS5nZXR0eWltYWdl/cy5jb20vaWQvODAy/ODkzNDI2L3Bob3Rv/L2x5Y2hlZS5qcGc_/cz02MTJ4NjEyJnc9/MCZrPTIwJmM9bVVX/dVZQYkstRENpZEF3/N3JSb2R5dU96a21t/UHpodFlzTHpkSE9L/Q2Ywaz0",
	}

	ProductList = append(ProductList, prd1)
	ProductList = append(ProductList, prd2)
	ProductList = append(ProductList, prd3)
	ProductList = append(ProductList, prd4)
	ProductList = append(ProductList, prd5)
	ProductList = append(ProductList, prd6)
}

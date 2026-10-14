
Chào bạn, hiện bạn dựa vào các câu lệnh này:
[
-- name: GetAllProductDetails :many
SELECT * FROM product_detail;

-- name: GetProductDetailByID :one
SELECT * FROM product_detail WHERE id_product = $1;

-- name: CreateProductDetail :exec
INSERT INTO product_detail(price, like_count, quantity, production_date, expiration_date) VALUES($1, $2, $3, $4, $5);

-- name: DeleteProductDetail :execresult
DELETE FROM product_detail WHERE id_product = $1;

-- name: UpdateProductDetailByID_PUT :execresult
UPDATE product_detail 
SET 	price = $1, 
		like_count = $2, 
		quantity = $3, 
		production_date = $4, 
		expiration_date = $5
WHERE id_product = $6;

-- name: UpdateProductPriceByID_PATCH :execresult
UPDATE product_detail 
SET 	price = $1
WHERE id_product = $2;

-- name: UpdateProductLikeCountByID_PATCH :execresult
UPDATE product_detail 
SET 	like_count = $1
WHERE id_product = $2;

-- name: UpdateProductQuantityByID_PATCH :execresult
UPDATE product_detail 
SET 	quantity = $1
WHERE id_product = $2;
]
Table đã tạo sẳn như này:[
create table PRODUCT_DETAIL (
   ID_PRODUCT      VARCHAR(50)    PRIMARY KEY,
   PRICE           NUMERIC(15, 2) NOT NULL,
   LIKE_COUNT      INT            NOT NULL DEFAULT 0,
   QUANTITY        INT            NOT NULL DEFAULT 0,
   PRODUCTION_DATE DATE,
   EXPIRATION_DATE DATE,
   UPDATED_AT      TIMESTAMP,
   CONSTRAINT FK_PRODUCT_DETAIL_PRODUCT
      FOREIGN KEY (ID_PRODUCT) REFERENCES PRODUCT (ID_PRODUCT)
         ON DELETE CASCADE ON UPDATE RESTRICT
);
]
Ở tệp tin 'sql\queries_product.sql'
Thì hiện tại tôi muốn bạn thực hiện trong phạm vi sau đây:
- Đầu tiên tạo tệp tin 'ProductDetail.go' lưu tại đây "internal\models" với tôi muốn nội dung bên trong lập tin là: "Định nghĩa struct của đối tượng này" ,bạn có thể tham khảo tệp tin này: "internal\models\status.go"

- thứ hai tạo tệp tin 'ProductDetail_mapper.go' lưu tại đây "internal\mapper" với tôi muốn nội dung bên trong lập tin là: "Xây dựng hàm map dữ liệu từ database.ProductDetail đã được tạo sẳn từ thư viện sqlc" ,bạn có thể tham khảo tệp tin này: "internal\mapper\status_mapper.go"

- thứ ba tạo tệp tin 'ProductDetail_repo.go' lưu tại đây "internal\repository\repository_impl" với tôi muốn nội dung bên trong lập tin là: "Chỉ định nghĩa Interface với các methobs bên trong sau cho tương ứng với cách hành động từ sql\queries_product.sql đã định nghĩa trước" ,bạn có thể tham khảo tệp tin này: "internal\repository\status_repo.go"

- thứ tư tạo tệp tin 'ProductDetail_repo_plm.go' lưu tại đây "internal\mapper" với tôi muốn nội dung bên trong lập tin là: triển khai Interface từ tệp tin ProductDetail_repo.go đã được định nghĩa trước đó" ,bạn có thể tham khảo tệp tin này: "internal\repository\repository_impl\status_repo_impl.go"

- thứ năm tạo tệp tin 'ProductDetail_service.go' lưu tại đây "internal\service" với tôi muốn nội dung bên trong lập tin là: "Vừa định nghĩa Interface ở tầng service này và vừa triển khai Interface tại đây. Và ở tầng service sẽ dựa vào tầng Repository chủ yếu qua Interface" ,bạn có thể tham khảo tệp tin này: "internal\service\status_service.go"

- thứ sáu tạo tệp tin 'ProductDetail.controller.go' lưu tại đây "internal\controller\http" với tôi muốn nội dung bên trong lập tin là: "Xây dựng các endpoint ở tầng controller này sẽ xử lý các request từ client và nó sẽ phụ thuộc vào tệp tin ProductDetail_service.go ở tầng service thông qua Interface" ,bạn có thể tham khảo tệp tin này: "internal\controller\http\status.controller.go"

- thứ bảy tạo tệp tin 'ProductDetail.wire.go' lưu tại đây "internal\wire" với tôi muốn nội dung bên trong lập tin là: "Xây dựng các phần khởi tạo giữa các tầng từ repository, service, controller để đảm bào cho việc dependency injection thông qua wire " ,bạn có thể tham khảo tệp tin này: "internal\wire\status.wire.go"

- thứ tám tạo tệp tin 'ProductDetail_mapper.go' lưu tại đây "internal\mapper" với tôi muốn nội dung bên trong lập tin là: "Xây dựng hàm map dữ liệu từ database.ProductDetail đã được tạo sẳn từ thư viện sqlc" ,bạn có thể tham khảo tệp tin này: "internal\routers\manager\status.router.go"

- thứ chín tạo tệp tin 'ProductDetail.router.go' lưu tại đây "internal\routers\manager" với tôi muốn nội dung bên trong lập tin là: "Định nghĩa các route để có thể gọi api để tương tác" ,bạn có thể tham khảo tệp tin này: "internal\routers\manager\status.router.go"

. Thì bạn hãy thực hiện theo tôi yêu cầu trong phạm vi đó. Không được vượt quá giới hạn
## Section 110: LÀM CHỦ MYSQL - Nguyên tắc ACID, sự cô lập vs cơ chế trong transaction HIỂU SÂU

# Hoạt động cơ bản của ecommerce mysql

-- USER A

1. Check sản phẩm tồn kho (100 sản phẩm)
   UPDATE `shopdev_inventory` SET ...; (100 - 2 = 98)

2. Tạo order
   INSERT INTO `shopdev_order` VALUES (...); (+2 sản phẩm)

3. Thanh toán tiền mặt
   INSERT INTO `shopdev_payment` VALUES (...);
4. Push qua cho giao hàng nhanh
   INSERT INTO `shopdev_logistics` VALUES (...) ;

-- USER B

1. Check sản phẩm tồn kho (100 sản phẩm)
   UPDATE `shopdev_inventory` SET ...; (100 - 2 = 98)

2. Tạo order
   INSERT INTO `shopdev_order` VALUES (...); (+2 sản phẩm)

3. Thanh toán tiền mặt
   INSERT INTO `shopdev_payment` VALUES (...);
4. Push qua cho giao hàng nhanh
   INSERT INTO `shopdev_logistics` VALUES (...) ;

---

# ACID

1. Atomicity (Tính nguyên tử)
   -Là một nhóm tạo lên 1 giao dịch (transaction) đều được thực hiện thành công hoặc thất bại.
2. Consistency (Tính nhất quán)

- Dữ liệu phải luôn phải nhất quán, không bị lỗi.
  EX: Tài khoản ngân hàng không thể bị trừ tiền mà không được cộng tiền vào tài khoản khác.
  EX: 1. Trước tồn kho là 100, 2. mua 2 sản phẩm, 3. sau đó tồn kho phải là 98,
  --> nhất quán: 98 + 2 = trước tồn kho (100)

3. Isolation (Tính cô lập)
   --> Đẩm bảo rằng mỗi giao dịch không ảnh hưởng đến các giao dịch khác.

4. Durability (Tính bền vững)
   --> Dữ liệu đã được lưu trữ sẽ không bị mất, không bị hỏng.
   --> Tất cả dữ liệu thay đổi sẽ được ghi vào ổ đĩa trên database.

# Bắt đầu transaction

START TRANSACTION | BEGIN | BEGIN WORK; -- Bắt đầu 1 giao dịch

COMMIT; -- Gửi giao dịch đến database

ROLLBACK; -- Khôi phục lại trạng thái trước khi giao dịch được thực hiện

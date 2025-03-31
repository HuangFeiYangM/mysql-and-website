

在 MySQL 中，`COUNT` 是一个 **聚合函数**，用于统计记录数量或非空值的数量。以下是它的核心用法和常见场景的详细解析：

---

### **一、基本语法**
```sql
COUNT(expression)  -- 统计非 NULL 的记录数
COUNT(*)           -- 统计所有行数（包括 NULL 行）
```

---

### **二、常见用法**

#### **1. 统计所有行数**
```sql
-- 统计 employee 表的总行数（无论是否有 NULL 值）
SELECT COUNT(*) AS Total_Employees FROM employee;
```
**输出示例**：
```
Total_Employees
7
```

#### **2. 统计某列非空值的数量**
```sql
-- 统计 Wno（仓库编号）列的非空值数量
SELECT COUNT(Wno) AS Non_Null_Wno FROM employee;
```
**示例表数据**：
假设 `Wno` 列有 2 个 `NULL` 值（如员工 0023 和 0031）
**输出**：
```
Non_Null_Wno
5
```

#### **3. 结合 `DISTINCT` 去重统计**
```sql
-- 统计仓库编号（Wno）的不同值数量
SELECT COUNT(DISTINCT Wno) AS Unique_Warehouses FROM employee;
```
**示例表数据**：
`Wno` 列有 `A01`, `A02`, `NULL`
**输出**：
```
Unique_Warehouses
2  -- 去除了 NULL 和重复值
```

---

### **三、进阶用法**

#### **1. 结合 `GROUP BY` 分组统计**
```sql
-- 按仓库统计员工人数（NULL 仓库单独分组）
SELECT Wno, COUNT(*) AS Employee_Count 
FROM employee 
GROUP BY Wno;
```
**输出示例**：
```
Wno   Employee_Count
A01   3
A02   3
NULL  1
```

#### **2. 结合 `CASE` 条件统计**
```sql
-- 统计薪资超过 3000 的员工数量
SELECT 
  COUNT(CASE WHEN Salary > 3000 THEN 1 ELSE NULL END) AS High_Salary_Count
FROM employee;
```
**输出示例**：
```
High_Salary_Count
4  -- 例如 张三(3600)、张立(8500)、王强(5600)、李军(5000)
```

---

### **四、`COUNT` 与 `NULL` 的陷阱**

- **`COUNT(*)` vs `COUNT(列名)`**
  - `COUNT(*)`：统计所有行，包括 `NULL`。
  - `COUNT(列名)`：统计该列非 `NULL` 的行数。

**示例**：
```sql
SELECT 
  COUNT(*) AS Total_Rows,
  COUNT(Wno) AS Non_Null_Wno,
  COUNT(DISTINCT Wno) AS Unique_Wno
FROM employee;
```
**输出**：
```
Total_Rows  Non_Null_Wno  Unique_Wno
7           5             2
```

---

### **五、性能优化**

#### **1. `COUNT(*)` 和 `COUNT(1)` 的区别**
在 MySQL 中，`COUNT(*)` 和 `COUNT(1)` 性能完全相同，优先用 `COUNT(*)`（语义更明确）。

#### **2. 快速统计大表**
如果表非常大，直接 `SELECT COUNT(*)` 可能很慢。
**优化方法**：
- 使用近似值（适合对精度要求不高的场景）：
  ```sql
  SHOW TABLE STATUS LIKE 'employee';  -- 查看 Rows 列的估算值
  ```
- 维护一个统计表（手动或通过触发器更新）。

---

### **六、实战案例**

#### **场景**
统计每个仓库的员工人数，并显示仓库名称（假设仓库信息在 `warehouse` 表中）。

**表结构**：
```sql
CREATE TABLE warehouse (
  Wno VARCHAR(10) PRIMARY KEY,
  Wname VARCHAR(30)
);
```

**SQL 代码**：
```sql
SELECT 
  w.Wno, 
  w.Wname, 
  COUNT(e.Eid) AS Employee_Count
FROM warehouse w
LEFT JOIN employee e ON w.Wno = e.Wno
GROUP BY w.Wno, w.Wname;
```

**输出示例**：
```
Wno  Wname    Employee_Count
A01  上海仓    3
A02  北京仓    3
```

---

### **七、常见错误**
1. **忘记 `GROUP BY`**
   ```sql
   -- 错误：直接统计分组后的数量
   SELECT Wno, COUNT(*) FROM employee;
   ```
   **修正**：必须使用 `GROUP BY Wno`。

2. **混淆 `COUNT` 和 `SUM`**
   ```sql
   -- 错误：用 SUM 统计行数
   SELECT SUM(1) FROM employee;  -- 正确但语义不清晰
   ```
   **建议**：明确使用 `COUNT(*)`。

---

掌握 `COUNT` 的用法后，你可以轻松应对大部分统计需求！如果有具体问题，欢迎继续提问 😊





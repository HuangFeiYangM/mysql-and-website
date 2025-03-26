
### **一、表结构定义**
#### **1. 学生表 `Student`**
```sql
CREATe TABLE Student(
    Sno   CHAR(5) primary key,       -- 学号，主键
    Sname CHAR(20) UNIQUE,           -- 姓名，唯一约束
    Ssex  CHAR(1) CHECK (ssex IN ('F','M')),  -- 性别，只能为 F/M
    Sage  INT,                        -- 年龄
    Sdept CHAR(15)                   -- 院系
);
```
- **注意**：`CREATe` 拼写不规范，建议统一使用 `CREATE`。

---

#### **2. 课程表 `course` 和 `course1`**
```sql
-- 课程表 course（Cno 为主键）
CREATE TABLE course(
    Cno   CHAR(5) PRIMARY KEY,
    Cname CHAR(20) UNIQUE
);

-- 课程表 course1（与 course 结构相同，但主键定义方式不同）
CREATE TABLE course1(
    Cno   CHAR(5),
    Cname CHAR(20) UNIQUE,
    PRIMARY KEY (cno)  -- 显式定义主键
);
```
- **问题**：`course1` 表与 `course` 冗余，无实际意义。

---

#### **3. 选课表 `SC` 和错误表 `SC1`**
```sql
-- 选课表 SC（复合主键 + 外键）
CREATE TABLE SC(
    Sno   CHAR(5),
    Cno   CHAR(5),
    Grade INT,
    PRIMARY KEY (Sno, Cno),                    -- 复合主键
    FOREIGN KEY (sno) REFERENCES Student(sno),  -- 外键引用 Student
    FOREIGN KEY (cno) REFERENCES course(cno)    -- 外键引用 course
);

-- 错误尝试：重复定义 SC1 表（第一次错误）
CREATE TABLE SC1(
    Sno CHAR(5) Primary key,  -- 错误：单列主键
    Cno CHAR(5) Primary key,   -- 错误：重复定义主键
    Grade INT,
    FOREIGN KEY (sno) REFERENCES Student(sno),
    FOREIGN KEY (cno) REFERENCES course(cno)
);

-- 修正后的 SC1（第二次定义）
CREATE TABLE SC1(
    Sno CHAR(5),
    Cno CHAR(5),
    Grade INT,
    PRIMARY KEY (Sno, Cno),                   -- 正确：复合主键
    FOREIGN KEY (sno) REFERENCES Student(sno),
    FOREIGN KEY (cno) REFERENCES course(cno)
);
```
- **错误**：
  - 第一次定义 `SC1` 时重复使用 `Primary key`，导致语法错误。
  - 多次创建同名表 `SC1`，需先删除旧表。

---

### **二、数据操作**
#### **1. 删除表与修改表结构**
```sql
DROP TABLE Student2;  -- 删除不存在的表（原代码中有 Student2，此处可能误删）
ALTER TABLE sc ADD COLUMN fcredit FLOAT;  -- 添加浮点类型列
```

#### **2. 插入数据**
```sql
-- 正确插入（省略 INTO 可能不兼容某些数据库）
INSERT INTO student(sno,sname) VALUES ('3301','zs');
INSERT student(sno,sname) VALUES ('3302','LS');        -- 缺少 INTO（建议补充）
INSERT INTO student(sno,sname,sage) 
VALUES ('33031',NULL,19), ('33042','',20);  -- 插入 NULL 和空字符串
```

#### **3. 复制表结构**
```sql
CREATE TABLE stu AS SELECT * FROM student;   -- 复制表结构及数据
CREATE TABLE stu1 SELECT * FROM student;     -- 同上（省略 AS）
```

---

### **三、索引操作**
```sql
-- 创建单列索引
CREATE INDEX wjlindexsname ON student(sname);

-- 创建复合索引（desc 在索引中无效，部分数据库不支持）
CREATE INDEX wjlindexsnameage ON student(sname, sage desc);
```
- **注意**：索引中的 `desc` 通常无效，排序需在查询时指定。

---

### **四、查询操作**
#### **1. 基础查询**
```sql
-- 错误：NULL 判断应用 IS NULL
SELECT * FROM student WHERE sage=19 AND sname=null;  -- 错误写法
SELECT * FROM student WHERE sage=19 AND sname IS NULL;  -- 正确写法

-- 空字符串查询
SELECT * FROM student WHERE sage=19 AND sname='';  -- 查询空字符串

-- 字符串函数
SELECT sno, LEFT(sno,2), sname FROM student ORDER BY sno DESC;  -- 截取前2字符
SELECT sno, SUBSTR(sno,3,3), sname FROM student ORDER BY sno DESC; -- 截取第3位起3字符
```

#### **2. 随机查询与分页**
```sql
-- 随机返回1条记录
SELECT * FROM student ORDER BY rand() LIMIT 1;

-- 分页查询（跳过2条，取3条）
SELECT * FROM student ORDER BY rand() LIMIT 2,3;
```

---

#### **3. 聚合与分组**
```sql
-- 统计总数
SELECT COUNT(*) AS num FROM student;

-- 分组统计（错误：sname 未在 GROUP BY 中）
SELECT ssex, sname, COUNT(*) 
FROM stu 
WHERE ssex='F' 
GROUP BY ssex, sname  -- 需包含 sname
HAVING COUNT(*) >1;
```

---

#### **4. 连接查询**
```sql
-- 隐式连接（不推荐）
SELECT Student.*, SC.*, course.* 
FROM Student, SC, course 
WHERE Student.Sno = SC.Sno AND course.cno=sc.cno;

-- 显式连接（推荐）
SELECT s.sno, sname, course.cno, cname, grade
FROM Student AS s 
JOIN SC ON s.Sno = SC.Sno 
JOIN course ON course.cno = sc.cno;

-- 右连接（保留所有学生，即使未选课）
SELECT Student.Sno, Sname, Ssex, Sage, Sdept, Cno, Grade
FROM SC 
RIGHT JOIN Student ON Student.Sno = SC.Sno;
```

---

### **五、常见错误总结**
1. **主键冲突**：单表不能定义多个主键。
2. **外键引用**：需确保引用的表和字段存在。
3. **NULL 处理**：必须用 `IS NULL` 或 `IS NOT NULL` 判断。
4. **隐式连接**：建议改用 `JOIN` 语法提高可读性。
5. **索引排序**：索引定义中的 `DESC` 通常无效。
6. **分组字段**：`SELECT` 中的非聚合字段需在 `GROUP BY` 中声明。

---

### **六、优化建议**
1. **统一命名规范**：表名、字段名统一大小写（如全小写）。
2. **使用显式 `INSERT INTO`**：避免省略关键字。
3. **避免冗余表**：如 `course1` 和 `SC1` 无必要存在。
4. **清理无效代码**：如重复的 `DROP TABLE Student2`。

通过修正这些问题，代码将更健壮且符合 SQL 标准。
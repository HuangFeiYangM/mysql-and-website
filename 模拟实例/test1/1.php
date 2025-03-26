<<<<<<< HEAD
header("Access-Control-Allow-Origin: *");           // 允许所有域名（生产环境应限制）
header("Access-Control-Allow-Methods: POST, GET");  // 允许的请求方法
header("Access-Control-Allow-Headers: Content-Type"); // 允许的请求头

error_reporting(E_ALL);
ini_set('display_errors', 1);



<?php
// 获取前端提交的数据
$username = $_POST['username'];
$email = $_POST['email'];

// 连接数据库
$servername = "192.168.31.210";
$db_username = "test1";
$db_password = "123456";
$dbname = "test1";

$conn = new mysqli($servername, $db_username, $db_password, $dbname);

// 检查连接
if ($conn->connect_error) {
    die("连接失败: " . $conn->connect_error);
}

// 使用预处理语句防止 SQL 注入
$stmt = $conn->prepare("INSERT INTO users (username, email) VALUES (?, ?)");
$stmt->bind_param("ss", $username, $email);

// 执行插入
if ($stmt->execute()) {
    echo "数据插入成功";
} else {
    echo "错误: " . $stmt->error;
}

// 关闭连接
$stmt->close();
$conn->close();
?>
=======
header("Access-Control-Allow-Origin: *");           // 允许所有域名（生产环境应限制）
header("Access-Control-Allow-Methods: POST, GET");  // 允许的请求方法
header("Access-Control-Allow-Headers: Content-Type"); // 允许的请求头

error_reporting(E_ALL);
ini_set('display_errors', 1);



<?php
// 获取前端提交的数据
$username = $_POST['username'];
$email = $_POST['email'];

// 连接数据库
$servername = "192.168.31.210";
$db_username = "test1";
$db_password = "123456";
$dbname = "test1";

$conn = new mysqli($servername, $db_username, $db_password, $dbname);

// 检查连接
if ($conn->connect_error) {
    die("连接失败: " . $conn->connect_error);
}

// 使用预处理语句防止 SQL 注入
$stmt = $conn->prepare("INSERT INTO users (username, email) VALUES (?, ?)");
$stmt->bind_param("ss", $username, $email);

// 执行插入
if ($stmt->execute()) {
    echo "数据插入成功";
} else {
    echo "错误: " . $stmt->error;
}

// 关闭连接
$stmt->close();
$conn->close();
?>
>>>>>>> main_1

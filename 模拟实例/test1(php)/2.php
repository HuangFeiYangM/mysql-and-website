<<<<<<< HEAD
<?php
// 处理跨域 OPTIONS 预检请求
if ($_SERVER['REQUEST_METHOD'] == 'OPTIONS') {
    header("HTTP/1.1 200 OK");
    exit();
}

// 设置响应头
header("Access-Control-Allow-Origin: *");
header("Access-Control-Allow-Methods: POST, GET, OPTIONS");
header("Access-Control-Allow-Headers: Content-Type");

// 开启错误报告
error_reporting(E_ALL);
ini_set('display_errors', 1);

// 检查是否为 POST 请求
if ($_SERVER['REQUEST_METHOD'] !== 'POST') {
    http_response_code(405);
    die("错误：仅支持 POST 请求");
}

// 检查 POST 数据是否存在
if (!isset($_POST['username']) || !isset($_POST['email'])) {
    http_response_code(400);
    die("错误：缺少用户名或邮箱");
}

// 获取数据
$username = $_POST['username'];
$email = $_POST['email'];

// 数据库配置
$servername = "192.168.31.210";
$db_username = "test1";
$db_password = "123456";
$dbname = "test1";

// 连接数据库
$conn = new mysqli($servername, $db_username, $db_password, $dbname);
if ($conn->connect_error) {
    die("连接失败: " . $conn->connect_error);
}

// 准备并执行 SQL
$sql = "INSERT INTO users (username, email) VALUES (?, ?)";
$stmt = $conn->prepare($sql);
if (!$stmt) {
    die("预处理失败: " . $conn->error);
}

$stmt->bind_param("ss", $username, $email);
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
<?php
// 处理跨域 OPTIONS 预检请求
if ($_SERVER['REQUEST_METHOD'] == 'OPTIONS') {
    header("HTTP/1.1 200 OK");
    exit();
}

// 设置响应头
header("Access-Control-Allow-Origin: *");
header("Access-Control-Allow-Methods: POST, GET, OPTIONS");
header("Access-Control-Allow-Headers: Content-Type");

// 开启错误报告
error_reporting(E_ALL);
ini_set('display_errors', 1);

// 检查是否为 POST 请求
if ($_SERVER['REQUEST_METHOD'] !== 'POST') {
    http_response_code(405);
    die("错误：仅支持 POST 请求");
}

// 检查 POST 数据是否存在
if (!isset($_POST['username']) || !isset($_POST['email'])) {
    http_response_code(400);
    die("错误：缺少用户名或邮箱");
}

// 获取数据
$username = $_POST['username'];
$email = $_POST['email'];

// 数据库配置
$servername = "192.168.31.210";
$db_username = "test1";
$db_password = "123456";
$dbname = "test1";

// 连接数据库
$conn = new mysqli($servername, $db_username, $db_password, $dbname);
if ($conn->connect_error) {
    die("连接失败: " . $conn->connect_error);
}

// 准备并执行 SQL
$sql = "INSERT INTO users (username, email) VALUES (?, ?)";
$stmt = $conn->prepare($sql);
if (!$stmt) {
    die("预处理失败: " . $conn->error);
}

$stmt->bind_param("ss", $username, $email);
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

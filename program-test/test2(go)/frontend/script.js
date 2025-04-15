// frontend/script.js
document.addEventListener('DOMContentLoaded', () => {
    // 检查用户是否已登录
    checkAuthStatus();
    
    // 事件监听
    document.getElementById('login-btn').addEventListener('click', showLoginForm);
    document.getElementById('register-btn').addEventListener('click', showRegisterForm);
    document.getElementById('logout-btn').addEventListener('click', logout);
    
    // 关闭模态框
    document.querySelector('.close').addEventListener('click', () => {
        document.getElementById('modal').style.display = 'none';
    });
});

function checkAuthStatus() {
    const token = localStorage.getItem('token');
    if (token) {
        document.getElementById('login-btn').style.display = 'none';
        document.getElementById('register-btn').style.display = 'none';
        document.getElementById('logout-btn').style.display = 'block';
        loadUserVehicles();
    }
}

function showLoginForm() {
    document.getElementById('modal-body').innerHTML = `
        <h2>Login</h2>
        <form id="login-form">
            <div>
                <label>Username:</label>
                <input type="text" name="username" required>
            </div>
            <div>
                <label>Password:</label>
                <input type="password" name="password" required>
            </div>
            <button type="submit">Login</button>
        </form>
    `;
    
    document.getElementById('modal').style.display = 'block';
    
    document.getElementById('login-form').addEventListener('submit', (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);
        const data = Object.fromEntries(formData.entries());
        login(data);
    });
}

// 类似实现 register 和 vehicle 表单

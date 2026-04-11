// 詳細情報の表示/非表示切り替え
document.addEventListener("DOMContentLoaded", () => {
    const btn = document.getElementById("toggleBtn");
    const detail = document.getElementById("detail");

    btn.addEventListener("click", () => {
        const isHidden = detail.style.display === "none";
        detail.style.display = isHidden ? "block" : "none"; 
        btn.textContent = isHidden ? "詳細情報を隠す" : "詳細情報の表示";
    });
});
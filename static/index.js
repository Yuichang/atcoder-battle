// ボタンが押された時に、2人のユーザーが存在するかを確認させる
document.addEventListener("DOMContentLoaded", () => {
    const form = document.querySelector("form");
    const user1Input = document.querySelector('input[name="user1"]');
    const user2Input = document.querySelector('input[name="user2"]');

    form.addEventListener("submit", async (e) => {
        e.preventDefault();

        const user1 = user1Input.value;
        const user2 = user2Input.value;

        const ok1 = await checkUser(user1);
        const ok2 = await checkUser(user2);

        if (!ok1 || !ok2) {
            alert("存在しないユーザーが含まれています");
            return;
        }

        form.submit();
    });
});

async function checkUser(username) {
    try {
        const res = await fetch(`/check_user?username=${username}`);
        console.log("status:", res.status);

        const data = await res.json();
        console.log("data:", data);

        return data.ok;
    } catch (err) {
        console.log("fetch error:", err);
        return false;
    }
}
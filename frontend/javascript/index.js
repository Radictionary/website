if (localStorage.getItem("--light-color") != null)
    document.documentElement.style.setProperty(
        "--light-color",
        localStorage.getItem("--light-color")
    );
if (localStorage.getItem("--dark-color") != null)
    document.documentElement.style.setProperty(
        "--dark-color",
        localStorage.getItem("--dark-color")
    );
const triangleSvg = document.getElementById("triangles");
const titleSvg = document.getElementById("title");
const r = document.querySelector(":root");

const colors = ["blue", "green", "orange", "yellow"];
titleSvg.onclick = (e) => {
    const rando = () => colors[Math.floor(Math.random() * colors.length)];
    document.documentElement.style.cssText = `
        --dark-color: ${rando()};
        --light-color: ${rando()};
        `;
    localStorage.setItem("--light-color", `${rando()}`);
    localStorage.setItem("--dark-color", `${rando()}`);
};

// Espera a que el DOM esté listo
document.addEventListener("DOMContentLoaded", () => {
  const form = document.getElementById("habitForm");
  const nameInput = document.getElementById("name");
  const typeInput = document.getElementById("type");
  const list = document.getElementById("habitList");

  // Función para cargar los hábitos desde la API
  async function loadHabits() {
    const res = await fetch("/api/habits");
    const data = await res.json();
    list.innerHTML = ""; // limpiar lista

    if (data.length === 0) {
      list.innerHTML = "<li>No hay hábitos registrados aún 🌱</li>";
      return;
    }

    data.forEach((habit) => {
      const li = document.createElement("li");
      li.classList.add("habit-item");
      li.innerHTML = `
        <span><b>${habit.name}</b> (${habit.type}) - ${habit.points} pts</span>
      `;
      list.appendChild(li);
    });
  }

  // Cuando se envía el formulario
  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const name = nameInput.value.trim();
    const type = typeInput.value.trim();

    if (!name || !type) {
      alert("Por favor, completa todos los campos 🌍");
      return;
    }

    await fetch("/api/habits", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name, type, points: 10 }),
    });

    nameInput.value = "";
    typeInput.value = "";
    await loadHabits();
  });

  // Cargar hábitos al inicio
  loadHabits();
});

const themeToggle = document.getElementById('theme-toggle');
const userTheme = localStorage.getItem('theme');

// Function to set theme based on preference
function applyTheme(theme) {
  if (theme === 'dark') {
    document.body.classList.remove('light-theme');
    document.body.classList.add('dark-theme');
  } else {
    document.body.classList.remove('dark-theme');
    document.body.classList.add('light-theme');
  }
}

// Check for saved user preference; otherwise, check OS theme
if (userTheme) {
  applyTheme(userTheme);
} else {
  const systemPrefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
  applyTheme(systemPrefersDark ? 'dark' : 'light');
}

// Listen for OS theme changes
window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
  if (!localStorage.getItem('theme')) {  // Only update if there's no saved user preference
    applyTheme(e.matches ? 'dark' : 'light');
  }
});

// Toggle theme manually with button click
themeToggle.addEventListener('click', () => {
  const newTheme = document.body.classList.contains('dark-theme') ? 'light' : 'dark';
  applyTheme(newTheme);
  localStorage.setItem('theme', newTheme);  // Save user preference
});

// Mobile menu toggle
document.addEventListener('DOMContentLoaded', () => {
  const mobileMenuBtn = document.getElementById('mobile-menu-btn');
  const mobileMenu = document.getElementById('mobile-menu');

  if (mobileMenuBtn && mobileMenu) {
    mobileMenuBtn.addEventListener('click', () => {
      const isExpanded = mobileMenuBtn.getAttribute('aria-expanded') === 'true';
      mobileMenuBtn.setAttribute('aria-expanded', !isExpanded);
      mobileMenu.classList.toggle('hidden');
    });
  }

  // Contact form submission
  const contactForm = document.getElementById('contact-form');
  if (contactForm) {
    contactForm.addEventListener('submit', async (e) => {
      e.preventDefault();

      const formData = new FormData(contactForm);
      const data = Object.fromEntries(formData.entries());

      // Check honeypot
      if (data.website) {
        console.log('Spam detected');
        return;
      }

      try {
        const response = await fetch('/contact', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(data),
        });

        if (response.ok) {
          alert('Thank you! We will be in touch soon.');
          contactForm.reset();
        } else {
          alert('There was an error submitting your message. Please try again.');
        }
      } catch (error) {
        console.error('Error:', error);
        alert('There was an error submitting your message. Please try again.');
      }
    });
  }

  // Work filter enhancements (optional progressive enhancement)
  const workFilters = document.querySelector('form[action="/work"]');
  if (workFilters) {
    // Could add instant filter updates via JS here
    // For now, relies on form submission (no-JS fallback)
  }
});

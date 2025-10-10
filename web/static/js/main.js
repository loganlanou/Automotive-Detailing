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

      // Show loading state
      const submitBtn = contactForm.querySelector('button[type="submit"]');
      const originalText = submitBtn.textContent;
      submitBtn.disabled = true;
      submitBtn.textContent = 'Sending...';

      try {
        const response = await fetch('/contact', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(data),
        });

        if (response.ok) {
          // Show success message
          const successMsg = document.createElement('div');
          successMsg.className = 'bg-green-500/10 border border-green-500/20 text-green-400 px-4 py-3 rounded-lg mb-4';
          successMsg.innerHTML = '<p class="font-semibold">Thank you for contacting us!</p><p class="text-sm mt-1">We\'ll get back to you within 2-4 business hours.</p>';
          contactForm.insertAdjacentElement('beforebegin', successMsg);
          contactForm.reset();
          setTimeout(() => successMsg.remove(), 5000);
        } else {
          throw new Error('Server error');
        }
      } catch (error) {
        console.error('Error:', error);
        // Show error message
        const errorMsg = document.createElement('div');
        errorMsg.className = 'bg-red-500/10 border border-red-500/20 text-red-400 px-4 py-3 rounded-lg mb-4';
        errorMsg.innerHTML = '<p class="font-semibold">Oops! Something went wrong.</p><p class="text-sm mt-1">Please try again or call us directly at (555) 123-4567.</p>';
        contactForm.insertAdjacentElement('beforebegin', errorMsg);
        setTimeout(() => errorMsg.remove(), 5000);
      } finally {
        // Reset button state
        submitBtn.disabled = false;
        submitBtn.textContent = originalText;
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

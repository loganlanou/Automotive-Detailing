// ========================================
// Intersection Observer for Scroll Animations
// ========================================
function initScrollAnimations() {
  const observerOptions = {
    threshold: 0.1,
    rootMargin: '0px 0px -50px 0px'
  };

  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('animate-in');
        // Optionally unobserve after animation
        observer.unobserve(entry.target);
      }
    });
  }, observerOptions);

  // Observe elements with fade-in, slide-up, scale-in classes
  document.querySelectorAll('.fade-in, .slide-up, .scale-in').forEach(el => {
    observer.observe(el);
  });
}

// ========================================
// Animated Stats Counter
// ========================================
function animateCounter(element, target, duration = 2000) {
  const start = 0;
  const increment = target / (duration / 16);
  let current = start;

  const timer = setInterval(() => {
    current += increment;
    if (current >= target) {
      element.textContent = target;
      clearInterval(timer);
    } else {
      // Handle decimal values
      if (target % 1 !== 0) {
        element.textContent = current.toFixed(1);
      } else {
        element.textContent = Math.floor(current);
      }
    }
  }, 16);
}

function initStatsCounter() {
  const statsSection = document.querySelector('.stats-section');
  if (!statsSection) return;

  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        const stats = entry.target.querySelectorAll('[data-count]');
        stats.forEach(stat => {
          const target = parseFloat(stat.dataset.count);
          const suffix = stat.dataset.suffix || '';
          const prefix = stat.dataset.prefix || '';

          animateCounter(stat, target, 2000);

          // Add suffix after animation
          setTimeout(() => {
            stat.textContent = prefix + stat.textContent + suffix;
          }, 2000);
        });
        observer.unobserve(entry.target);
      }
    });
  }, { threshold: 0.5 });

  observer.observe(statsSection);
}

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

  // Initialize scroll animations
  initScrollAnimations();

  // Initialize stats counter
  initStatsCounter();

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

  // Initialize image lightbox
  initLightbox();

  // Initialize before/after sliders
  initBeforeAfterSliders();

  // Initialize testimonials carousel
  initTestimonialsCarousel();

  // Initialize pricing calculator
  initPricingCalculator();

  // Work filter enhancements (optional progressive enhancement)
  const workFilters = document.querySelector('form[action="/work"]');
  if (workFilters) {
    // Could add instant filter updates via JS here
    // For now, relies on form submission (no-JS fallback)
  }
});

// ========================================
// Image Lightbox
// ========================================
function initLightbox() {
  const lightboxTriggers = document.querySelectorAll('[data-lightbox]');
  if (lightboxTriggers.length === 0) return;

  // Create lightbox container
  const lightbox = document.createElement('div');
  lightbox.id = 'lightbox';
  lightbox.className = 'fixed inset-0 bg-black/90 z-50 hidden items-center justify-center p-4';
  lightbox.innerHTML = `
    <button class="absolute top-4 right-4 text-white hover:text-brand-accent transition p-2" aria-label="Close lightbox">
      <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
      </svg>
    </button>
    <button class="absolute left-4 top-1/2 -translate-y-1/2 text-white hover:text-brand-accent transition p-2" aria-label="Previous image">
      <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
      </svg>
    </button>
    <button class="absolute right-4 top-1/2 -translate-y-1/2 text-white hover:text-brand-accent transition p-2" aria-label="Next image">
      <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
      </svg>
    </button>
    <img src="" alt="" class="max-w-full max-h-full object-contain rounded-lg">
  `;
  document.body.appendChild(lightbox);

  const closeBtn = lightbox.querySelector('button[aria-label="Close lightbox"]');
  const prevBtn = lightbox.querySelector('button[aria-label="Previous image"]');
  const nextBtn = lightbox.querySelector('button[aria-label="Next image"]');
  const img = lightbox.querySelector('img');

  let currentIndex = 0;
  const images = Array.from(lightboxTriggers);

  function showImage(index) {
    currentIndex = index;
    const trigger = images[index];
    img.src = trigger.dataset.lightbox || trigger.src;
    img.alt = trigger.alt || '';
    lightbox.classList.remove('hidden');
    lightbox.classList.add('flex');
    document.body.style.overflow = 'hidden';
  }

  function closeLightbox() {
    lightbox.classList.add('hidden');
    lightbox.classList.remove('flex');
    document.body.style.overflow = '';
  }

  function showNext() {
    currentIndex = (currentIndex + 1) % images.length;
    showImage(currentIndex);
  }

  function showPrev() {
    currentIndex = (currentIndex - 1 + images.length) % images.length;
    showImage(currentIndex);
  }

  lightboxTriggers.forEach((trigger, index) => {
    trigger.addEventListener('click', (e) => {
      e.preventDefault();
      showImage(index);
    });
  });

  closeBtn.addEventListener('click', closeLightbox);
  nextBtn.addEventListener('click', showNext);
  prevBtn.addEventListener('click', showPrev);

  lightbox.addEventListener('click', (e) => {
    if (e.target === lightbox) closeLightbox();
  });

  document.addEventListener('keydown', (e) => {
    if (!lightbox.classList.contains('hidden')) {
      if (e.key === 'Escape') closeLightbox();
      if (e.key === 'ArrowRight') showNext();
      if (e.key === 'ArrowLeft') showPrev();
    }
  });
}

// ========================================
// Before/After Image Slider
// ========================================
function initBeforeAfterSliders() {
  const sliders = document.querySelectorAll('.before-after-slider');

  sliders.forEach(slider => {
    const beforeImg = slider.querySelector('.before-image');
    const afterImg = slider.querySelector('.after-image');
    const handle = slider.querySelector('.slider-handle');

    if (!beforeImg || !afterImg || !handle) return;

    let isDragging = false;

    function updateSlider(percentage) {
      percentage = Math.max(0, Math.min(100, percentage));
      afterImg.style.clipPath = `inset(0 ${100 - percentage}% 0 0)`;
      handle.style.left = `${percentage}%`;
    }

    function handleMove(e) {
      if (!isDragging) return;

      const rect = slider.getBoundingClientRect();
      const x = (e.type.includes('touch') ? e.touches[0].clientX : e.clientX) - rect.left;
      const percentage = (x / rect.width) * 100;

      updateSlider(percentage);
    }

    handle.addEventListener('mousedown', () => isDragging = true);
    handle.addEventListener('touchstart', () => isDragging = true);

    document.addEventListener('mouseup', () => isDragging = false);
    document.addEventListener('touchend', () => isDragging = false);

    document.addEventListener('mousemove', handleMove);
    document.addEventListener('touchmove', handleMove);

    // Initialize at 50%
    updateSlider(50);
  });
}

// ========================================
// Testimonials Carousel
// ========================================
function initTestimonialsCarousel() {
  const carousel = document.querySelector('.testimonials-carousel');
  if (!carousel) return;

  const track = carousel.querySelector('.carousel-track');
  const slides = carousel.querySelectorAll('.carousel-slide');
  const prevBtn = carousel.querySelector('.carousel-prev');
  const nextBtn = carousel.querySelector('.carousel-next');
  const dotsContainer = carousel.querySelector('.carousel-dots');

  if (!track || slides.length === 0) return;

  let currentSlide = 0;
  const totalSlides = slides.length;

  // Create dots
  if (dotsContainer) {
    for (let i = 0; i < totalSlides; i++) {
      const dot = document.createElement('button');
      dot.className = 'w-2 h-2 rounded-full bg-muted hover:bg-brand-accent transition';
      dot.setAttribute('aria-label', `Go to slide ${i + 1}`);
      dot.addEventListener('click', () => goToSlide(i));
      dotsContainer.appendChild(dot);
    }
  }

  function updateCarousel() {
    track.style.transform = `translateX(-${currentSlide * 100}%)`;

    // Update dots
    if (dotsContainer) {
      const dots = dotsContainer.querySelectorAll('button');
      dots.forEach((dot, index) => {
        if (index === currentSlide) {
          dot.classList.add('bg-brand-accent');
          dot.classList.remove('bg-muted');
        } else {
          dot.classList.remove('bg-brand-accent');
          dot.classList.add('bg-muted');
        }
      });
    }
  }

  function goToSlide(index) {
    currentSlide = index;
    updateCarousel();
  }

  function nextSlide() {
    currentSlide = (currentSlide + 1) % totalSlides;
    updateCarousel();
  }

  function prevSlide() {
    currentSlide = (currentSlide - 1 + totalSlides) % totalSlides;
    updateCarousel();
  }

  if (nextBtn) nextBtn.addEventListener('click', nextSlide);
  if (prevBtn) prevBtn.addEventListener('click', prevSlide);

  // Auto-rotate every 5 seconds
  setInterval(nextSlide, 5000);

  updateCarousel();
}

// ========================================
// Pricing Calculator
// ========================================
function initPricingCalculator() {
  const calculator = document.getElementById('pricing-calculator');
  if (!calculator) return;

  const vehicleType = calculator.querySelector('[name="vehicle-type"]');
  const servicePackage = calculator.querySelector('[name="service-package"]');
  const addons = calculator.querySelectorAll('[name="addon"]');
  const priceDisplay = calculator.querySelector('.price-display');

  const prices = {
    sedan: { exterior: 200, interior: 200, full: 400 },
    suv: { exterior: 250, interior: 250, full: 500 },
    truck: { exterior: 300, interior: 300, full: 600 },
    luxury: { exterior: 350, interior: 350, full: 700 }
  };

  const addonPrices = {
    ceramic: 300,
    paint: 150,
    engine: 100,
    headlight: 75
  };

  function calculatePrice() {
    if (!vehicleType || !servicePackage || !priceDisplay) return;

    const vehicle = vehicleType.value;
    const service = servicePackage.value;

    let basePrice = prices[vehicle]?.[service] || 0;
    let addonPrice = 0;

    addons.forEach(addon => {
      if (addon.checked) {
        addonPrice += addonPrices[addon.value] || 0;
      }
    });

    const total = basePrice + addonPrice;
    priceDisplay.textContent = `$${total}`;
  }

  if (vehicleType) vehicleType.addEventListener('change', calculatePrice);
  if (servicePackage) servicePackage.addEventListener('change', calculatePrice);
  addons.forEach(addon => addon.addEventListener('change', calculatePrice));

  // Initial calculation
  calculatePrice();
}

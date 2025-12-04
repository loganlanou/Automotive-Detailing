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

// Mobile menu toggle with smooth animation
document.addEventListener('DOMContentLoaded', () => {
  const mobileMenuBtn = document.getElementById('mobile-menu-btn');
  const mobileMenu = document.getElementById('mobile-menu');

  if (mobileMenuBtn && mobileMenu) {
    mobileMenuBtn.addEventListener('click', () => {
      const isExpanded = mobileMenuBtn.getAttribute('aria-expanded') === 'true';
      mobileMenuBtn.setAttribute('aria-expanded', !isExpanded);
      mobileMenu.classList.toggle('show');

      // Animate hamburger icon
      const svg = mobileMenuBtn.querySelector('svg');
      if (isExpanded) {
        svg.style.transform = 'rotate(0deg)';
      } else {
        svg.style.transform = 'rotate(90deg)';
      }
    });
  }

  // Header scroll effect
  const header = document.getElementById('main-header');
  let lastScroll = 0;

  window.addEventListener('scroll', () => {
    const currentScroll = window.scrollY;

    if (currentScroll > 50) {
      header.classList.add('header-scrolled');
    } else {
      header.classList.remove('header-scrolled');
    }

    lastScroll = currentScroll;
  }, { passive: true });

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

      // Check honeypot (spam protection)
      if (data.website) {
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

  // Initialize floating CTA
  initFloatingCTA();
});

// ========================================
// Floating CTA Button (appears on scroll)
// ========================================
function initFloatingCTA() {
  const floatingCTA = document.querySelector('.floating-cta');
  if (!floatingCTA) return;

  let ticking = false;

  function updateCTA() {
    const scrollY = window.scrollY;

    // Show CTA after scrolling down 300px or past hero section
    if (scrollY > 300) {
      floatingCTA.classList.add('visible');
    } else {
      floatingCTA.classList.remove('visible');
    }
    ticking = false;
  }

  function requestTick() {
    if (!ticking) {
      window.requestAnimationFrame(updateCTA);
      ticking = true;
    }
  }

  window.addEventListener('scroll', requestTick, { passive: true });

  // Initial check
  updateCTA();
}

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
    const afterContainer = slider.querySelector('#after-container');
    const afterImg = slider.querySelector('#after-image');
    const handle = slider.querySelector('#slider-handle');

    if (!afterContainer || !afterImg || !handle) return;

    let isDragging = false;

    function updateSlider(clientX) {
      const rect = slider.getBoundingClientRect();
      const x = clientX - rect.left;
      const percentage = (x / rect.width) * 100;
      const boundedPercentage = Math.max(0, Math.min(100, percentage));

      // Update the width of the after container
      afterContainer.style.width = `${boundedPercentage}%`;
      // Adjust the after image to maintain full width
      afterImg.style.width = `${rect.width}px`;
      // Position the handle
      handle.style.left = `${boundedPercentage}%`;
    }

    function handleMove(e) {
      if (!isDragging) return;
      e.preventDefault();

      const clientX = e.type.includes('touch') ? e.touches[0].clientX : e.clientX;
      updateSlider(clientX);
    }

    function startDrag(e) {
      isDragging = true;
      e.preventDefault();
    }

    function stopDrag() {
      isDragging = false;
    }

    // Mouse events
    handle.addEventListener('mousedown', startDrag);
    slider.addEventListener('mousedown', (e) => {
      isDragging = true;
      updateSlider(e.clientX);
    });
    document.addEventListener('mousemove', handleMove);
    document.addEventListener('mouseup', stopDrag);

    // Touch events
    handle.addEventListener('touchstart', startDrag);
    slider.addEventListener('touchstart', (e) => {
      isDragging = true;
      updateSlider(e.touches[0].clientX);
    });
    document.addEventListener('touchmove', handleMove, { passive: false });
    document.addEventListener('touchend', stopDrag);

    // Initialize at 50%
    window.addEventListener('load', () => {
      const rect = slider.getBoundingClientRect();
      updateSlider(rect.left + rect.width / 2);
    });
  });
}

// ========================================
// Social Sharing Functions
// ========================================
function shareOnFacebook() {
  const url = encodeURIComponent(window.location.href);
  window.open(`https://www.facebook.com/sharer/sharer.php?u=${url}`, '_blank', 'width=600,height=400');
}

function shareOnTwitter() {
  const url = encodeURIComponent(window.location.href);
  const title = encodeURIComponent(document.title);
  window.open(`https://twitter.com/intent/tweet?url=${url}&text=${title}`, '_blank', 'width=600,height=400');
}

function copyLink() {
  const url = window.location.href;

  if (navigator.clipboard && window.isSecureContext) {
    navigator.clipboard.writeText(url).then(() => {
      showToast('Link copied to clipboard!');
    }).catch(() => {
      fallbackCopyLink(url);
    });
  } else {
    fallbackCopyLink(url);
  }
}

function fallbackCopyLink(text) {
  const textArea = document.createElement('textarea');
  textArea.value = text;
  textArea.style.position = 'fixed';
  textArea.style.left = '-999999px';
  document.body.appendChild(textArea);
  textArea.focus();
  textArea.select();

  try {
    document.execCommand('copy');
    showToast('Link copied to clipboard!');
  } catch (err) {
    showToast('Failed to copy link');
  }

  document.body.removeChild(textArea);
}

function showToast(message) {
  const toast = document.createElement('div');
  toast.className = 'fixed bottom-4 right-4 bg-brand-accent text-white px-6 py-3 rounded-lg shadow-lg z-50 animate-in fade-in slide-up';
  toast.textContent = message;
  document.body.appendChild(toast);

  setTimeout(() => {
    toast.classList.add('opacity-0', 'transition-opacity');
    setTimeout(() => toast.remove(), 300);
  }, 3000);
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

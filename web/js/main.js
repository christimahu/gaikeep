/**
 * main.js - Core functionality for Gai Keep website
 * Contains event handling, animations, and UI interactions
 */
(() => {
  'use strict';

  /**
   * Initialize all website functionality when DOM is fully loaded
   */
  document.addEventListener('DOMContentLoaded', () => {
    handleSplashScreen();
    initAnimations();
    initNavigation();
    initCardEffects();
    initTouchInteractions();
  });

  /**
   * Handles splash screen display and dismissal
   */
  const handleSplashScreen = () => {
    const splashScreen = document.getElementById('splash-screen');
    
    // Hide splash screen after content loads
    if (splashScreen) {
      // If in standalone mode (added to home screen), show splash longer
      const isInStandaloneMode = window.matchMedia('(display-mode: standalone)').matches || 
                               (window.navigator.standalone || false);
      
      // Longer timeout for first visit or when launched from home screen
      const timeoutDuration = isInStandaloneMode ? 2000 : 1000;
      
      setTimeout(() => {
        splashScreen.classList.add('splash-hidden');
        setTimeout(() => {
          splashScreen.style.display = 'none';
        }, 500); // Match the CSS transition duration
      }, timeoutDuration);
    }
  };

  /**
   * Manages section animations based on viewport visibility
   */
  const initAnimations = () => {
    const sections = document.querySelectorAll('section');
    
    /**
     * Determines if an element is sufficiently in the viewport to animate
     * @param {HTMLElement} element - The element to check
     * @returns {boolean} - Whether element should be animated
     */
    const isInViewport = (element) => {
      const rect = element.getBoundingClientRect();
      const offset = 100; // pixels from bottom of viewport to trigger animation
      return rect.top < window.innerHeight - offset;
    };

    /**
     * Adds visibility class to elements in viewport
     */
    const animateVisibleSections = () => {
      sections.forEach(section => {
        if (isInViewport(section)) {
          section.classList.add('visible');
        }
      });
    };

    // Initial check on page load
    animateVisibleSections();
    
    // Check again on scroll, using throttling for performance
    let scrollTimeout;
    window.addEventListener('scroll', () => {
      if (!scrollTimeout) {
        scrollTimeout = setTimeout(() => {
          animateVisibleSections();
          scrollTimeout = null;
        }, 100);
      }
    });
  };

  /**
   * Sets up all navigation-related functionality
   */
  const initNavigation = () => {
    const navLinks = document.querySelectorAll('.navigation a');
    const sections = document.querySelectorAll('section');
    
    /**
     * Implements smooth scrolling for navigation links
     */
    const setupSmoothScroll = () => {
      navLinks.forEach(link => {
        link.addEventListener('click', (event) => {
          event.preventDefault();
          
          const targetId = link.getAttribute('href');
          const targetElement = document.querySelector(targetId);
          
          if (targetElement) {
            window.scrollTo({
              top: targetElement.offsetTop,
              behavior: 'smooth'
            });
            
            // Update URL without page reload
            history.pushState(null, '', targetId);
            
            // Close mobile menu if open
            closeMobileMenuIfOpen();
          }
        });
      });
    };
    
    /**
     * Closes mobile menu when a link is clicked
     */
    const closeMobileMenuIfOpen = () => {
      const mobileMenu = document.getElementById('mobile-menu');
      if (mobileMenu && !mobileMenu.classList.contains('hidden') && window.innerWidth < 768) {
        mobileMenu.classList.add('hidden');
      }
    };
    
    /**
     * Highlights the active navigation link based on scroll position
     */
    const setupActiveNavigation = () => {
      // Use IntersectionObserver for better performance
      const observerOptions = {
        root: null,
        rootMargin: '-20% 0px -70%',
        threshold: 0
      };
      
      const observerCallback = (entries) => {
        entries.forEach(entry => {
          if (entry.isIntersecting) {
            const currentId = entry.target.getAttribute('id');
            
            // Remove active class from all links
            navLinks.forEach(link => {
              link.classList.remove('active');
            });
            
            // Add active class to current link
            const currentLink = document.querySelector(`.navigation a[href="#${currentId}"]`);
            if (currentLink) {
              currentLink.classList.add('active');
            }
          }
        });
      };
      
      const observer = new IntersectionObserver(observerCallback, observerOptions);
      
      // Observe all sections
      sections.forEach(section => {
        observer.observe(section);
      });
    };
    
    setupSmoothScroll();
    setupActiveNavigation();
  };
  
  /**
   * Initializes interactive effects for card elements
   */
  const initCardEffects = () => {
    const cards = document.querySelectorAll('.card-hover');
    
    cards.forEach(card => {
      // Add shadow and transform effects on hover
      card.addEventListener('mouseenter', () => {
        card.classList.add('shadow-lg');
        card.style.transform = 'translateY(-5px)';
      });
      
      card.addEventListener('mouseleave', () => {
        card.classList.remove('shadow-lg');
        card.style.transform = 'translateY(0)';
      });
      
      // Add focus accessibility for keyboard navigation
      card.addEventListener('focus', () => {
        card.classList.add('shadow-lg');
      });
      
      card.addEventListener('blur', () => {
        card.classList.remove('shadow-lg');
      });
    });
  };
  
  /**
   * Initializes touch-specific interactions for mobile devices
   */
  const initTouchInteractions = () => {
    // Add touch feedback for buttons and links
    const touchElements = document.querySelectorAll('a, button, .card-hover');
    
    touchElements.forEach(element => {
      // Add active state on touch
      element.addEventListener('touchstart', () => {
        element.classList.add('touch-active');
      }, { passive: true });
      
      // Remove active state after touch
      ['touchend', 'touchcancel'].forEach(eventType => {
        element.addEventListener(eventType, () => {
          element.classList.remove('touch-active');
        }, { passive: true });
      });
    });
    
    // Improve iOS scrolling
    document.documentElement.style.webkitTouchCallout = 'none';
    
    // Handle viewport height issues on iOS (addressing the "100vh problem")
    const setViewportHeight = () => {
      const vh = window.innerHeight * 0.01;
      document.documentElement.style.setProperty('--vh', `${vh}px`);
    };
    
    // Set initial viewport height
    setViewportHeight();
    
    // Update on resize and orientation change
    window.addEventListener('resize', () => {
      setViewportHeight();
    });
    
    window.addEventListener('orientationchange', () => {
      setTimeout(setViewportHeight, 100);
    });
  };
  
  /**
   * Toggles the mobile menu visibility
   * This function is called directly from HTML
   * @global
   */
  window.toggleMenu = () => {
    const mobileMenu = document.getElementById('mobile-menu');
    if (mobileMenu) {
      mobileMenu.classList.toggle('hidden');
    }
  };
})();

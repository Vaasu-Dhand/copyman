export const updateFavicon = () => {
    // Remove existing favicon links
    const existingLinks = document.querySelectorAll('link[rel*="icon"]');
    existingLinks.forEach(link => link.remove());

    // Create new favicon link
    const favicon = document.createElement('link');
    favicon.rel = 'icon';
    favicon.type = 'image/png';
    favicon.href = '/favicon-light.png';

    // Add to head
    document.head.appendChild(favicon);

    // Also update apple touch icon if you have them
    const appleTouchIcon = document.createElement('link');
    appleTouchIcon.rel = 'apple-touch-icon';
    appleTouchIcon.sizes = '180x180';
    appleTouchIcon.href = '/apple-touch-icon-light.png';

    document.head.appendChild(appleTouchIcon);
};
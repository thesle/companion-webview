// Wait for Wails runtime to be ready
window.addEventListener('DOMContentLoaded', async () => {
    const iframe = document.getElementById('webview-frame');
    const loading = document.getElementById('loading');
    
    try {
        // Get the URL from the Go backend
        const url = await window.go.main.App.GetURL();
        
        // Set the iframe source
        iframe.src = url;
        
        // Hide loading message once iframe loads
        iframe.addEventListener('load', () => {
            loading.style.display = 'none';
        });
        
        // Handle iframe load errors
        iframe.addEventListener('error', () => {
            loading.textContent = 'Error loading URL: ' + url;
        });
        
    } catch (error) {
        loading.textContent = 'Error: ' + error.message;
        console.error('Error loading URL:', error);
    }
    
    // Setup keyboard shortcuts - use capture phase to intercept before iframe
    const handleKeyboard = async (e) => {
        // Ctrl+W to close window
        if (e.ctrlKey && e.key === 'w') {
            e.preventDefault();
            e.stopPropagation();
            await window.go.main.App.CloseWindow();
            return false;
        }
        
        // Ctrl+M to maximize/restore window
        if (e.ctrlKey && e.key === 'm') {
            e.preventDefault();
            e.stopPropagation();
            await window.go.main.App.MaximizeWindow();
            return false;
        }
        
        // Ctrl+Shift+M to minimize window
        if (e.ctrlKey && e.shiftKey && e.key === 'M') {
            e.preventDefault();
            e.stopPropagation();
            await window.go.main.App.MinimizeWindow();
            return false;
        }
    };
    
    // Add listener in capture phase to intercept before iframe
    document.addEventListener('keydown', handleKeyboard, true);
    window.addEventListener('keydown', handleKeyboard, true);
});

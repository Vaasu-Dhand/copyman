import React, { useState, useEffect } from 'react';
// import { GetSettings, SaveSettings, CopyToClipboard, HideOverlay } from '../wailsjs/go/main/App';
import './App.css';

interface Settings {
  theme: string;
  keyBindings: { [key: string]: string };
}

const App: React.FC = () => {
  const [settings, setSettings] = useState<Settings>({
    theme: 'light',
    keyBindings: {
      '1': '', '2': '', '3': '', '4': '', '5': '', 
      '6': '', '7': '', '8': '', '9': ''
    }
  });
  
  const [currentView, setCurrentView] = useState<'main' | 'settings'>('main');
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    loadSettings();
  }, []);

  useEffect(() => {
    document.documentElement.setAttribute('data-theme', settings.theme);
  }, [settings.theme]);

  useEffect(() => {
    // Handle keyboard events
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.key === 'Escape') {
        HideOverlay();
      }
      
      // Handle number keys 1-9
      if (event.key >= '1' && event.key <= '9') {
        const text = settings.keyBindings[event.key];
        if (text) {
          handleCopyText(text);
        }
      }
    };

    window.addEventListener('keydown', handleKeyDown);
    return () => window.removeEventListener('keydown', handleKeyDown);
  }, [settings.keyBindings]);

  const loadSettings = async () => {
    try {
      const loadedSettings = await GetSettings();
      setSettings(loadedSettings);
    } catch (error) {
      console.error('Failed to load settings:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const saveSettings = async (newSettings: Settings) => {
    try {
      await SaveSettings(newSettings);
      setSettings(newSettings);
    } catch (error) {
      console.error('Failed to save settings:', error);
    }
  };

  const handleCopyText = async (text: string) => {
    if (!text) return;
    
    try {
      await CopyToClipboard(text);
    } catch (error) {
      console.error('Failed to copy to clipboard:', error);
    }
  };

  const handleKeyBinding = (key: string, value: string) => {
    const newSettings = {
      ...settings,
      keyBindings: {
        ...settings.keyBindings,
        [key]: value
      }
    };
    saveSettings(newSettings);
  };

  const toggleTheme = () => {
    const newTheme = settings.theme === 'light' ? 'dark' : 'light';
    const newSettings = {
      ...settings,
      theme: newTheme
    };
    saveSettings(newSettings);
  };

  if (isLoading) {
    return (
      <div className="app-container">
        <div className="loading">Loading...</div>
      </div>
    );
  }

  const MainView = () => (
    <div className="overlay">
      <div className="keyboard-container">
        <div className="keyboard-header">
          <div className="keyboard-title">CopyMan</div>
          <div className="keyboard-subtitle">Press a number to copy</div>
        </div>
        
        <div className="keyboard-grid">
          {[1, 2, 3, 4, 5, 6, 7, 8, 9].map(num => (
            <div 
              key={num} 
              className="key" 
              onClick={() => handleCopyText(settings.keyBindings[num.toString()])}
            >
              <div className="key-number">{num}</div>
              <div className="key-text">
                {settings.keyBindings[num.toString()] || (
                  <span className="empty-state">Empty</span>
                )}
              </div>
            </div>
          ))}
        </div>
        
        <div className="keyboard-footer">
          <button 
            className="settings-button"
            onClick={() => setCurrentView('settings')}
          >
            ⚙️ Settings
          </button>
          <div className="close-hint">
            Press Esc to close • ⌘⇧C to open
          </div>
        </div>
      </div>
    </div>
  );

  const SettingsView = () => (
    <div className="settings-container">
      <div className="settings-header">
        <button 
          className="back-button"
          onClick={() => setCurrentView('main')}
        >
          ← Back
        </button>
        <div className="settings-title">Settings</div>
        <div className="settings-subtitle">Customize your clipboard shortcuts</div>
      </div>
      
      <div className="theme-toggle">
        <div className="toggle-header">
          <div className="toggle-title">Dark Mode</div>
          <div 
            className={`toggle-switch ${settings.theme === 'dark' ? 'active' : ''}`}
            onClick={toggleTheme}
          >
            <div className="toggle-knob"></div>
          </div>
        </div>
        <div className="toggle-description">
          Switch between light and dark appearance
        </div>
      </div>
      
      <div className="bindings-section">
        <div className="bindings-title">Key Bindings</div>
        <div className="bindings-grid">
          {[1, 2, 3, 4, 5, 6, 7, 8, 9].map(num => (
            <div key={num} className="binding-item">
              <div className="binding-number">{num}</div>
              <input
                type="text"
                className="binding-input"
                placeholder={`Text for key ${num}...`}
                value={settings.keyBindings[num.toString()]}
                onChange={(e) => handleKeyBinding(num.toString(), e.target.value)}
              />
            </div>
          ))}
        </div>
      </div>
    </div>
  );

  return (
    <div className="app-container">
      {currentView === 'main' ? <MainView /> : <SettingsView />}
    </div>
  );
};

export default App;
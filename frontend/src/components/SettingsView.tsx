import React, { useState, useEffect } from 'react';
import { Settings } from '../App';

interface SettingsViewProps {
  settings: Settings;
  onSettingsChange: (newSettings: Settings) => void;
  onNavigateBack: () => void;
}

const SettingsView: React.FC<SettingsViewProps> = ({ 
  settings, 
  onSettingsChange, 
  onNavigateBack 
}) => {
  // Local state for input values to prevent re-render focus issues
  const [localKeyBindings, setLocalKeyBindings] = useState<{ [key: string]: string }>({});

  // Initialize local state when settings change
  useEffect(() => {
    setLocalKeyBindings(settings.keyBindings);
  }, [settings.keyBindings]);

  const handleInputChange = (key: string, value: string) => {
    setLocalKeyBindings(prev => ({
      ...prev,
      [key]: value
    }));
  };

  const handleInputBlur = (key: string) => {
    const value = localKeyBindings[key];
    const newSettings = {
      ...settings,
      keyBindings: {
        ...settings.keyBindings,
        [key]: value
      }
    };
    onSettingsChange(newSettings);
  };

  const toggleTheme = () => {
    const newTheme = settings.theme === 'light' ? 'dark' : 'light';
    const newSettings = {
      ...settings,
      theme: newTheme
    };
    onSettingsChange(newSettings);
  };

  return (
    <div className="settings-container">
      <div className="settings-header">
        <button 
          className="back-button"
          onClick={onNavigateBack}
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
                value={localKeyBindings[num.toString()] || ''}
                onChange={(e) => handleInputChange(num.toString(), e.target.value)}
                onBlur={() => handleInputBlur(num.toString())}
              />
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default SettingsView;
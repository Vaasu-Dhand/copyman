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
    <div className="settings-container-horizontal">
      {/* Header with back button and theme toggle in same row */}
      <div className="settings-header-horizontal">
        <div className="header-left">
          <button 
            className="back-button-horizontal"
            onClick={onNavigateBack}
          >
            ← Back
          </button>
          <div className="title-group">
            <div className="settings-title-horizontal">Settings</div>
            <div className="settings-subtitle-horizontal">Customize your clipboard shortcuts</div>
          </div>
        </div>
        
        <div className="theme-toggle-horizontal">
          <div className="toggle-label">Dark Mode</div>
          <div 
            className={`toggle-switch-horizontal ${settings.theme === 'dark' ? 'active' : ''}`}
            onClick={toggleTheme}
          >
            <div className="toggle-knob-horizontal"></div>
          </div>
        </div>
      </div>
      
      {/* Key bindings in a compact grid */}
      <div className="bindings-section-horizontal">
        <div className="bindings-grid-horizontal">
          {[1, 2, 3, 4, 5, 6, 7, 8, 9].map(num => (
            <div key={num} className="binding-item-horizontal">
              <div className="binding-number-horizontal">{num}</div>
              <input
                type="text"
                className="binding-input-horizontal"
                placeholder={`Key ${num} text...`}
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
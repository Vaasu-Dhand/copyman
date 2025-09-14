import React, { useState, useEffect } from 'react';
import { GetSettings, SaveSettings } from '../wailsjs/go/main/App';
import MainView from './components/MainView';
import SettingsView from './components/SettingsView';
import { updateFavicon } from './utils/favicon';
import './App.css';

export interface Settings {
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
    updateFavicon();
  }, [settings.theme]);

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

  if (isLoading) {
    return (
      <div className="app-container">
        <div className="loading">Loading...</div>
      </div>
    );
  }

  return (
    <div className="app-container">
      {currentView === 'main' ? (
        <MainView 
          settings={settings} 
          onNavigateToSettings={() => setCurrentView('settings')} 
        />
      ) : (
        <SettingsView 
          settings={settings} 
          onSettingsChange={saveSettings}
          onNavigateBack={() => setCurrentView('main')} 
        />
      )}
    </div>
  );
};

export default App;
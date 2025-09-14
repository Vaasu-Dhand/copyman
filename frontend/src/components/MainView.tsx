import React, { useEffect, useState } from 'react';
import { CopyToClipboard, HideOverlay } from '../../wailsjs/go/main/App';
import { Settings } from '../App';

interface MainViewProps {
    settings: Settings;
    onNavigateToSettings: () => void;
}

const MainView: React.FC<MainViewProps> = ({ settings, onNavigateToSettings }) => {
    const [highlightedKey, setHighlightedKey] = useState<string | null>(null);

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
                    handleCopyText(text, event.key);
                }
            }
        };

        window.addEventListener('keydown', handleKeyDown);
        return () => window.removeEventListener('keydown', handleKeyDown);
    }, [settings.keyBindings]);

    const handleCopyText = async (text: string, keyNumber?: string) => {
        if (!text) return;

        try {
            await CopyToClipboard(text);

            // Highlight the key briefly
            if (keyNumber) {
                setHighlightedKey(keyNumber);
                setTimeout(() => setHighlightedKey(null), 200);
            }
        } catch (error) {
            console.error('Failed to copy to clipboard:', error);
        }
    };

    return (
        <div className="overlay">
            <div className="keyboard-container">
                <div className="keyboard-header">
                    <div className="keyboard-title">
                        <img src="favicon-light.png" alt="logo" width={20} height={20} />
                    </div>
                    <div className="keyboard-subtitle">Press a number to copy</div>
                </div>

                <div className="keyboard-grid">
                    {/* First row: 1-5 */}
                    <div className="keyboard-row">
                        {[1, 2, 3, 4, 5].map(num => {
                            const numStr = num.toString();
                            return (
                                <div
                                    key={num}
                                    className={`key ${highlightedKey === numStr ? 'highlighted' : ''}`}
                                    onClick={() => handleCopyText(settings.keyBindings[numStr], numStr)}
                                >
                                    <div className="key-number">{num}</div>
                                    <div className="key-text">
                                        {settings.keyBindings[numStr] || (
                                            <span className="empty-state">Empty</span>
                                        )}
                                    </div>
                                </div>
                            );
                        })}
                    </div>

                    {/* Second row: 6-9 */}
                    <div className="keyboard-row">
                        {[6, 7, 8, 9].map(num => {
                            const numStr = num.toString();
                            return (
                                <div
                                    key={num}
                                    className={`key ${highlightedKey === numStr ? 'highlighted' : ''}`}
                                    onClick={() => handleCopyText(settings.keyBindings[numStr], numStr)}
                                >
                                    <div className="key-number">{num}</div>
                                    <div className="key-text">
                                        {settings.keyBindings[numStr] || (
                                            <span className="empty-state">Empty</span>
                                        )}
                                    </div>
                                </div>
                            );
                        })}
                    </div>
                </div>

                <div className="keyboard-footer">
                    <button
                        className="settings-button"
                        onClick={onNavigateToSettings}
                    >
                        ⚙️ Settings
                    </button>
                    <div className="close-hint">
                        Press Esc to close • ⌃⇧C to open
                    </div>
                </div>
            </div>
        </div>
    );
};

export default MainView;
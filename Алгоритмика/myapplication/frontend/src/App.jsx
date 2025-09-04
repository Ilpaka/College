import React, { useState, useEffect } from 'react';
import { 
    GenerateSlice, 
    BubbleSort, 
    ParallelBatcherSort,
    QuickSort,
    TournamentSort,
    BinaryInsertionSort
} from "../wailsjs/go/main/App";
import './App.css';

function App() {
    const [array, setArray] = useState([]);
    const [sortedArray, setSortedArray] = useState([]);
    const [selectedAlgorithm, setSelectedAlgorithm] = useState('bubble');
    const [arraySize, setArraySize] = useState(50);
    const [isGenerating, setIsGenerating] = useState(false);
    const [isSorting, setIsSorting] = useState(false);
    const [sortTime, setSortTime] = useState(0);
    const [resultText, setResultText] = useState("Генерируйте массив для начала работы");

    const algorithms = [
        { id: 'bubble', name: '🫧 Пузырьковая сортировка', color: '#ff6b6b' },
        { id: 'quick', name: '⚡ Быстрая сортировка', color: '#4ecdc4' },
        { id: 'batcher', name: '🔄 Параллельная Бэтчера', color: '#45b7d1' },
        { id: 'tournament', name: '🏆 Турнирная сортировка', color: '#f9ca24' },
        { id: 'binary', name: '🎯 Бинарные вставки', color: '#6c5ce7' }
    ];

    const generateArray = async () => {
        setIsGenerating(true);
        setResultText("Генерация массива...");
        
        try {
            const newArray = await GenerateSlice(arraySize);
            setArray(newArray);
            setSortedArray([]);
            setSortTime(0);
            setResultText(`Сгенерирован массив из ${arraySize} элементов`);
        } catch (error) {
            setResultText("Ошибка генерации массива");
        }
        
        setIsGenerating(false);
    };

    const sortArray = async () => {
        if (array.length === 0) {
            setResultText("Сначала сгенерируйте массив!");
            return;
        }

        setIsSorting(true);
        const startTime = performance.now();
        setResultText("Сортировка в процессе...");

        try {
            let result = [];
            const selectedAlgo = algorithms.find(algo => algo.id === selectedAlgorithm);

            switch (selectedAlgorithm) {
                case 'bubble':
                    result = await BubbleSort(array);
                    break;
                case 'quick':
                    result = await QuickSort(array);
                    break;
                case 'batcher':
                    result = await ParallelBatcherSort(array);
                    break;
                case 'tournament':
                    result = await TournamentSort(array);
                    break;
                case 'binary':
                    result = await BinaryInsertionSort(array);
                    break;
                default:
                    result = await BubbleSort(array);
            }

            const endTime = performance.now();
            const duration = Math.round((endTime - startTime) * 100) / 100;

            setSortedArray(result);
            setSortTime(duration);
            setResultText(`${selectedAlgo.name} завершена за ${duration}мс`);
        } catch (error) {
            setResultText("Ошибка при сортировке");
        }

        setIsSorting(false);
    };

    const getArrayDisplay = (arr) => {
        if (arr.length === 0) return "[ ]";
        if (arr.length > 20) {
            return `[${arr.slice(0, 10).join(', ')}, ..., ${arr.slice(-10).join(', ')}]`;
        }
        return `[${arr.join(', ')}]`;
    };

    useEffect(() => {
        generateArray();
    }, []);

    return (
        <div id="App">
            <div className="neon-container">
                <header className="app-header">
                    <h1 className="neon-title">
                        <span className="neon-text">SORT</span>
                        <span className="neon-accent">MASTER</span>
                    </h1>
                    <p className="subtitle">Визуализатор алгоритмов сортировки</p>
                </header>

                <div className="main-content">
                    <div className="control-panel">
                        <div className="control-group">
                            <label className="control-label">Размер массива</label>
                            <div className="size-controls">
                                <input 
                                    type="range" 
                                    min="10" 
                                    max="200" 
                                    value={arraySize}
                                    onChange={(e) => setArraySize(parseInt(e.target.value))}
                                    className="neon-slider"
                                />
                                <span className="size-display">{arraySize}</span>
                            </div>
                        </div>

                        <div className="control-group">
                            <label className="control-label">Алгоритм сортировки</label>
                            <div className="algorithm-selector">
                                {algorithms.map(algo => (
                                    <div 
                                        key={algo.id}
                                        className={`algorithm-card ${selectedAlgorithm === algo.id ? 'selected' : ''}`}
                                        style={{
                                            '--accent-color': algo.color,
                                            '--accent-glow': algo.color + '40'
                                        }}
                                        onClick={() => setSelectedAlgorithm(algo.id)}
                                    >
                                        <span className="algo-name">{algo.name}</span>
                                    </div>
                                ))}
                            </div>
                        </div>

                        <div className="action-buttons">
                            <button 
                                className={`neon-button generate ${isGenerating ? 'loading' : ''}`}
                                onClick={generateArray}
                                disabled={isGenerating}
                            >
                                {isGenerating ? 'Генерация...' : '🎲 Генерировать массив'}
                            </button>

                            <button 
                                className={`neon-button sort ${isSorting ? 'loading' : ''}`}
                                onClick={sortArray}
                                disabled={isSorting || array.length === 0}
                            >
                                {isSorting ? 'Сортировка...' : '🚀 Сортировать'}
                            </button>
                        </div>
                    </div>

                    <div className="results-section">
                        <div className="status-display">
                            <div className="status-text">{resultText}</div>
                            {sortTime > 0 && (
                                <div className="performance-stats">
                                    <span className="stat">⏱️ Время: {sortTime}мс</span>
                                    <span className="stat">📊 Элементов: {array.length}</span>
                                </div>
                            )}
                        </div>

                        <div className="arrays-display">
                            {array.length > 0 && (
                                <div className="array-container">
                                    <h3 className="array-title original">Исходный массив</h3>
                                    <div className="array-content">
                                        {getArrayDisplay(array)}
                                    </div>
                                </div>
                            )}

                            {sortedArray.length > 0 && (
                                <div className="array-container">
                                    <h3 className="array-title sorted">Отсортированный массив</h3>
                                    <div className="array-content">
                                        {getArrayDisplay(sortedArray)}
                                    </div>
                                </div>
                            )}
                        </div>
                    </div>
                </div>

                <div className="floating-particles">
                    {[...Array(20)].map((_, i) => (
                        <div key={i} className={`particle particle-${i % 4}`}></div>
                    ))}
                </div>
            </div>
        </div>
    );
}

export default App;

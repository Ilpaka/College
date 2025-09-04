import { useState } from 'react';
import { 
  GenerateSlice, 
  BubbleSort, 
  QuickSort, 
  ParallelBatcherSort, 
  TournamentSort, 
  BinaryInsertionSort 
} from '../wailsjs/go/main/App';

function App() {
  const [originalArray, setOriginalArray] = useState([]);
  const [sortedArray, setSortedArray] = useState([]);
  const [selectedSort, setSelectedSort] = useState('bubble');
  const [arraySize, setArraySize] = useState(100);
  const [executionTime, setExecutionTime] = useState(null);


  // Доступные алгоритмы
  const sortingMethods = [
    { value: 'bubble', label: 'Пузырьковая сортировка', func: BubbleSort },
    { value: 'quick', label: 'Быстрая сортировка', func: QuickSort },
    { value: 'batcher', label: 'Параллельная сортировка Бэтчера', func: ParallelBatcherSort },
    { value: 'tournament', label: 'Турнирная сортировка', func: TournamentSort },
    { value: 'binary', label: 'Бинарная сортировка вставками', func: BinaryInsertionSort }
  ];

  // Генерация массива
  const generateArray = async () => {
    const newArray = await GenerateSlice(arraySize);
    setOriginalArray(newArray);
    setSortedArray([]); // Очищаем отсортированный массив
    setExecutionTime(null);
  };

  // Сортировка массива
  const sortArray = async () => {
    if (originalArray.length === 0) {
      alert('Сначала сгенерируйте массив!');
      return;
    }

    const selectedMethod = sortingMethods.find(method => method.value === selectedSort);
    const startTime = performance.now();
    const result = await selectedMethod.func(originalArray);
    const endTime = performance.now();
    const timeTaken = endTime - startTime;
    setSortedArray(result);
    setExecutionTime(timeTaken);
  };

  return (
    <div>
      <h1>Алгоритмы сортировки</h1>
      
      <div>
        <label>
          Размер массива: 
          <input
            type="number"
            value={arraySize}
            onChange={(e) => setArraySize(parseInt(e.target.value) || 100)}
            min="10"
            max="1000"
          />
        </label>
      </div>

      <div>
        <label>
          Алгоритм сортировки:
          <select 
            value={selectedSort} 
            onChange={(e) => setSelectedSort(e.target.value)}
          >
            {sortingMethods.map(method => (
              <option key={method.value} value={method.value}>
                {method.label}
              </option>
            ))}
          </select>
        </label>
      </div>

      <div>
        <button onClick={generateArray}>
          Сгенерировать массив
        </button>
        
        <button onClick={sortArray}>
          Отсортировать
        </button>
      </div>

    {executionTime !== null && (
    <div>
          <h3>⏱️ Время выполнения: {executionTime.toFixed(2)} мс</h3>
          <p>Алгоритм: {sortingMethods.find(m => m.value === selectedSort)?.label}</p>
        </div>
      )}
      <div>
        <h3>Исходный массив:</h3>
        <textarea
          readOnly
          value={originalArray.length > 0 ? originalArray.join(', ') : 'Массив не сгенерирован'}
          rows="5"
          cols="80"
        />
      </div>

      <div>
        <h3>Отсортированный массив:</h3>
        <textarea
          readOnly
          value={sortedArray.length > 0 ? sortedArray.join(', ') : 'Массив не отсортирован'}
          rows="5"
          cols="80"
        />
      </div>
    </div>
  );
}

export default App;

import { useState } from 'react';

import { 
  GenerateSlice, 
  BubbleSort, 
  QuickSort, 
  ShellSort,
  InsertionSort
} from '../wailsjs/go/main/App.js';

/*************  ✨ Windsurf Command ⭐  *************/
  /**
   * @description
   * App - это React-компонент, который тестирует различные алгоритмы
   * сортировки.
   *
   * @param {Object} props - props-ы, которые передаются из родительского
   * компонента.
   *
   * @return {ReactElement} - React-элемент, который отображает интерфейс
   * для тестирования алгоритмов.
   */
/*******  444192c7-ae4b-443b-aabb-a95f5db4fed3  *******/
function App() {
  const [slice, setSlice] = useState([]);
  const [sorted, setSorted] = useState([]);
  const [executionTime, setExecutionTime] = useState('');
  const [selectedMethod, setSelectedMethod] = useState('quick');

  const sortingMethods = [
    { value: 'bubble', label: 'Пузырьковая сортировка', func: BubbleSort },
    { value: 'quick', label: 'Быстрая сортировка', func: QuickSort },
    { value: 'shell', label: 'Сортировка шелла', func: ShellSort },
    { value: 'insertion', label: 'Сортировка выбором', func: InsertionSort },
  ];

  const generateSlice = async () => {
    const newSlice = await GenerateSlice(10000);
    setSlice(newSlice);
    setSorted([]);
    setExecutionTime('');
  };

  const sorter = async () => {
    if (slice.length === 0) {
      alert("Сгенерируйте массив");
      return;
    }
    
    const selectedMethodObj = sortingMethods.find(method => method.value === selectedMethod);
    
    const start = performance.now();
    const result = await selectedMethodObj.func(slice);
    const end = performance.now();
    
    setSorted(result);
    setExecutionTime(`${(end - start).toFixed(2)} ms`);
    console.log(`Время выполнения: ${(end - start).toFixed(2)} ms`);
  };

  return (
    <div>
      <h1>Тестируем алгоритмы сортировки</h1>

      <div>
        <label>Выберите алгоритм:
          <select value={selectedMethod} onChange={e => setSelectedMethod(e.target.value)}>
            {sortingMethods.map(method => (
              <option key={method.value} value={method.value}>
                {method.label}
              </option>
            ))}
          </select>
        </label>

        <div>
          <button onClick={generateSlice}>
            Сгенерировать
          </button>
          <button onClick={sorter}>
            Сортировать
          </button>
        </div>

        {executionTime && <h3>Время выполнения: {executionTime}</h3>}

        <div>
          <h3>Исходный массив: [{slice.join(', ')}]</h3>
          <h3>Отсортированный: [{sorted.join(', ')}]</h3>
        </div>
      </div>
    </div>
  );
}

export default App;

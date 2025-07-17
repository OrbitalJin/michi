import { useCallback, useState } from 'react'
import { Button } from './components/ui/button'
import { Input } from './components/ui/input';

function App() {
  const [count, setCount] = useState(0)

  const handleSubmit = useCallback(() => {
    window.location.href = 'https://google.com';
    setCount(count + 1);
  }, [count])

  return (
    <div className='flex h-screen items-center justify-center'>
      <div className='flex flex-row gap-2'>
        <Input
          placeholder='Search...'
          onSubmit={handleSubmit}
        />

        <Button onClick={handleSubmit}>
          Hello, {count}
        </Button>
      </div>
    </div>
  )
}

export default App

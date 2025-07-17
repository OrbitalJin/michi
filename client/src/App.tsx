import { useState } from 'react'
import { Button } from './components/ui/button'

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className='flex h-screen items-center justify-center bg-red-400'>
      <Button onClick={() => { setCount(count + 1); }}>
        Hello, {count}
      </Button>
    </div>
  )
}

export default App

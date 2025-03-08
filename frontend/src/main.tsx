import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import './index.css'
import App from './App.tsx'
import CreatePost from './pages/CreatePost'
import ViewPost from './pages/ViewPost'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="/posts/create" element={<CreatePost />} />
        <Route path="/posts/:id" element={<ViewPost />} />
      </Routes>
    </BrowserRouter>
  </StrictMode>,
)

import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { BrowserRouter } from 'react-router'
import { ThemeProvider } from './providers/themeProvider.tsx'
import { ConfigProvider } from './providers/configProvider.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <BrowserRouter>
      <ThemeProvider th='light'>
        <ConfigProvider>
          <App />
        </ConfigProvider>
      </ThemeProvider>
    </BrowserRouter>
  </StrictMode>,
)

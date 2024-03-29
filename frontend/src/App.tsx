import React, { lazy, Suspense } from 'react';
import { SnackbarProvider } from 'notistack';
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom';
import { AppContainer } from './styled-components';
import { Header } from '@/components';
import { topSecretContext } from '@/contexts';
import { useEffect, useState } from 'react';
import { SnackbarUtilsConfigurator } from './utilities';

import './App.scss'

export const myContext = topSecretContext
const TopSecretPage = lazy(() => import('@/pages/Alliance/TopSecretPage'));
const App = () => {
  const [satellite, setSatellite] = useState([]);
  return (
    <React.StrictMode>
        <Header 
            fontSizeTittle="1.1rem"
            valueTittle="Operacion Quasar"
        />
        <AppContainer className="App">
            <myContext.Provider value={{satellite, setSatellite}}>
            <SnackbarProvider>
             <SnackbarUtilsConfigurator />
              <Suspense fallback={<div>Loading ...</div>}>
                  <BrowserRouter>
                    <Routes>
                      <Route path="/" element={<Navigate to={`TopSecret`} />} />
                      <Route
                        path={`topSecret/*`}
                        element={
                          <div>
                             <TopSecretPage /> 
                          </div>
                        }
                      />
                    </Routes>
                  </BrowserRouter>
              </Suspense>
            </SnackbarProvider>
            </myContext.Provider>
        </AppContainer>
    </React.StrictMode>
  )
}

export default App

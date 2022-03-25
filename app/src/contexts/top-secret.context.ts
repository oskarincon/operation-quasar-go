import { createContext } from 'react';

export const topSecretContext = createContext({
  satellite: [],
  setSatellite: (user: any) => {}
});

import axios from 'axios';

export const postTopSecret = (params:any, headers: any) => {
  return axios.post<any>('http://127.0.0.1:3000/alliance/topsecret',
        params,
        headers
  );
};

import axios from 'axios';

export const postTopSecret = (params:any, headers: any) => {
  return axios.post<any>('https://0byvg60aae.execute-api.us-east-1.amazonaws.com/v1/alliance/topsecret',
        params,
        headers
  );
};

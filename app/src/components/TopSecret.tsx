import { Input } from '@/components';
import { postTopSecret } from '@/services';
import { Grid, Button } from '@mui/material';
import { InputType } from './Input';
import { useForm } from 'react-hook-form';
import { useContext, useState } from 'react';
import { topSecretContext } from '@/contexts';
import { Send } from '@mui/icons-material';

interface TopSecretProps {
  data: Array<any>;
  dataRequired?: Array<any>;
}


export const TopSecret = ({ data }: TopSecretProps) => {
  const { satellite, setSatellite } = useContext(topSecretContext)
  const [response, setResponse] = useState({
    mensajeResp: '',
    posicionX: '',
    posicionY: ''
  });
  const { register } = useForm();
  let inputVal: any;
  const dataRequired = data.map((evet)=>evet.name )

  const validateField = async (e: any) => {
    let aux = 0;
    dataRequired.map((evet)=>{
      const { id, type, value } = e.target;
      if(evet === id) {
          inputVal ={
            ...inputVal,
            [id+type]: value
          }
      }
    })
    for (const key in inputVal) {
      if (inputVal[key] !== '') {
        aux++;
      }
    }
  }

  const serviceCall = () => {
    const params = {
      satellites: [
        {
           name: "kenobi",
           distance: parseInt(inputVal.kenobinumber),
           message: inputVal.kenobitext.split()
        },
             {
           name: "skywalker",
           distance: parseInt(inputVal.skywalkernumber),
           message: inputVal.skywalkertext.split()
        },
             {
           name: "sato",
           distance: parseInt(inputVal.satonumber),
           message: inputVal.satotext.split()
        }
     ]
    }
    const headers = {
      'Content-Type': 'application/json'
    }
    const getApiData = async () => {
      try {
        const result = await postTopSecret(params, headers);
        adaptResponse(result.data);
      } catch (e) {}
    };
  
    const adaptResponse = (data: any) => {
      setResponse({
        mensajeResp: data.message,
        posicionX: String(data.position.x),
        posicionY: String(data.position.y)
      });
    };
  
    getApiData();
  };
  
  return (
  <>
   <Grid  container>
     <Grid container justifyContent="center" item xs={12} spacing={3}>
          {data.map((value) => (
            <Grid key={value.index} item>
                <p>{value.name}</p>
               <Input label="distance" name={value.name} type={InputType.NUMBER} register={register} trigger={async (e:any)=> await validateField(e)} ></Input>
               <br/>
               <Input label="mesagges" name={value.name} type={InputType.TEXT} register={register} trigger={async (e:any)=> await validateField(e)}  ></Input>
            </Grid>
          ))}
     </Grid>
   </Grid>
   <div style={{ alignContent: "center", padding: "1.5rem" }}>
     <Button variant="contained" color="primary"  endIcon={<Send />} onClick={() => { serviceCall(); }}>
       Validar
     </Button>
  </div>
  <p>Data Recibida</p>
  <Grid container justifyContent="center" item xs={12} >
    <div style={{ paddingRight: "0.1rem", width: "45%"}}>
        <p>Posicion X</p>
        <Input label="" name="X" type={InputType.TEXT} register={register} value={response.posicionX} disabled></Input>
    </div>
    <div style={{ paddingLeft: "0.1rem", width: "45%"}}>
        <p>Posicion Y</p>
        <Input label="" name="Y" type={InputType.TEXT} register={register}  disabled value={response.posicionY}></Input>
    </div>
  </Grid>
  <br/>
  <p>Mensaje</p>
  <Input label="" name="y" type={InputType.TEXT} register={register} value={response.mensajeResp}></Input>
  <br/>
    </>
  );
};

export default TopSecret;

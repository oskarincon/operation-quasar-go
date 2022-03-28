import { postTopSecret } from '@/services';
import { styled } from '@mui/material/styles';
import { Grid, Button, TextField, Tooltip, tooltipClasses, TooltipProps, InputAdornment } from '@mui/material';
import { useForm } from 'react-hook-form';
import { SnackbarUtilities } from '@/utilities';
import { useContext, useState, useEffect } from 'react';
import { topSecretContext } from '@/contexts';
import { Send, HelpTwoTone, AddLocation, Message } from '@mui/icons-material';
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";

interface TopSecretProps {
  data: Array<any>;
  dataRequired?: Array<any>;
}

const helpMessages = "Para la cadena de los mensajes se debe separar por comas; EJEMPLO: mensaje,de,,prueba,";

const CustomWidthTooltip = styled(({ className, ...props }: TooltipProps) => (
  <Tooltip {...props} classes={{ popper: className }} />
))({
  [`& .${tooltipClasses.tooltip}`]: {
    maxWidth: 500,
  },
});

const schema = yup.object().shape({
  kenobiDistance: yup.number().typeError("Distancia kenobi debe ser numerica").min(-999, "Distancia kenobi debe ser mayor a -999").max(999, "Distancia kenobi debe ser menor a 999").required(),
  skywalkerDistance: yup.number().typeError("Distancia skywalker debe ser numerica").min(-999, "Distancia skywalker debe ser mayor a -999").max(999, "Distancia skywalker debe ser menor a 999").required(),
  satoDistance: yup.number().typeError("Distancia sato debe ser numerica").min(-999, "Distancia sato debe ser mayor a -999").max(999, "Distancia sato debe ser menor a 999").required(),
  kenobiMessages: yup.string().min(5, "Mensaje kenobi debe ser mayor a 5 caracteres").max(50, "Mensaje kenobi debe ser menor a 50 caracteres").required(),
  skywalkerMessages: yup.string().min(5, "Mensaje skywalker debe ser mayor a 5 caracteres").max(50, "Mensaje skywalker debe ser menor a 50 caracteres").required(),
  satoMessages: yup.string().min(5, "Mensaje sato debe ser mayor a 5 caracteres").max(50, "Mensaje sato debe ser menor a 50 caracteres").required(),
})

export const TopSecret = ({ data }: TopSecretProps) => {
  const { satellite, setSatellite } = useContext(topSecretContext)
  const [response, setResponse] = useState({
    mensajeResp: '',
    posicionX: '',
    posicionY: ''
  });

  
  const { register, handleSubmit, setError, formState: { errors, isValid }   } = useForm({
    mode: "onChange",
    resolver: yupResolver(schema)
  });
  
  useEffect(() => {
    let errorSend = errors.kenobiDistance || errors.kenobiMessages || errors.skywalkerDistance || errors.skywalkerMessages || errors.satoDistance || errors.satoMessages
    if(errorSend?.message) {
      SnackbarUtilities.info(errorSend.message)
    }
  }, [errors.kenobiDistance || errors.kenobiMessages ||
      errors.skywalkerDistance || errors.skywalkerMessages ||
      errors.satoDistance || errors.satoMessages 
  ])

  const serviceCall = (params: any) => {
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
  
  const onSubmit = async (data: any) => {
    console.log(data);
    const { kenobiMessages, skywalkerMessages, satoMessages, satoDistance, skywalkerDistance, kenobiDistance } = data;
    if(!kenobiMessages.includes(',') || !skywalkerMessages.includes(',') || !satoMessages.includes(',')) {
      SnackbarUtilities.error('Se debe separar los mensajes con coma(,)')
      return
    }
    const params = {
      satellites: [
        {
           name: "kenobi",
           distance: parseInt(kenobiDistance),
           message: kenobiMessages.split(",")
        },
             {
           name: "skywalker",
           distance: parseInt(skywalkerDistance),
           message: skywalkerMessages.split(",")
        },
             {
           name: "sato",
           distance: parseInt(satoDistance),
           message: data.satoMessages.split(",")
        }
     ]
    }
    serviceCall(params);
  };
  
  return (
  <>
  <form onSubmit={handleSubmit(onSubmit)}>
     <Grid  container>
       <Grid container justifyContent="center" item xs={12} spacing={3}>
            {data.map((value) => (
              <Grid key={value.index} item>
                       <p>{value.name}  <span>
                 <CustomWidthTooltip title={helpMessages}>
                    <HelpTwoTone fontSize="medium" />
                </CustomWidthTooltip>
                 </span ></p>
                 <div>
                 <TextField 
                    label="Distance"
                    name={value.name + "Distance"}
                    type="number"
                    variant="outlined"
                    {...register(value.name + "Distance")} 
                  />
                 </div>
                 <br/>
                 <div>
                 <TextField 
                    label="Messages"
                    type="text"
                    name={value.name + "Messages"}
                    variant="outlined"
                    {...register(value.name + "Messages")} />
                 </div>
              </Grid>
            ))}
       </Grid>
     </Grid>
     <div style={{ alignContent: "center", padding: "1.5rem" }}>
       <Button variant="contained" color="primary" size="large" endIcon={<Send />} type="submit" id="submit" disabled={!isValid} >
         Validar
       </Button>
    </div>
  </form>
  <p>Data Response</p>
  <Grid container justifyContent="center" item xs={12} >
    <div style={{ paddingRight: "0.1rem", width: "45%"}}>
      <div>
      <TextField
        id="position-x"
        label="Posicion X"
        disabled
        value={response.posicionX}
        InputProps={{
          startAdornment: (
            <InputAdornment position="end">
              <AddLocation />
            </InputAdornment>
          ),
        }}
        variant="outlined"
      />
      </div>
    </div>
    <div style={{ paddingLeft: "0.1rem", width: "45%"}}>
    <div>
      <TextField
        id="position-y"
        label="Posicion Y"
        disabled
        value={response.posicionY}
        InputProps={{
          startAdornment: (
            <InputAdornment position="end">
              <AddLocation />
            </InputAdornment>
          ),
        }}
        variant="outlined"
      />
      </div>
    </div>
  </Grid>
  <br/>
  <div>
      <TextField
        id="messages-result"
        label="Messages Result"
        disabled
        value={response.mensajeResp}
        InputProps={{
          startAdornment: (
            <InputAdornment position="end">
              <Message />
            </InputAdornment>
          ),
        }}
        variant="outlined"
      />
      </div>
  <br/>
    </>
  );
};

export default TopSecret;

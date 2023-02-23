import React from 'react';
import { useDispatch, useSelector } from 'react-redux';
import Typography from '@mui/material/Typography';
import TextField from '@mui/material/TextField';
import Paper from '@mui/material/Paper';
import Button from '@mui/material/Button';
import Avatar from '@mui/material/Avatar';

import styles from './Login.module.scss';
import { fetchRegister, selectIsAuth } from "../../redux/slices/auth";
import { useForm } from 'react-hook-form';
import { Navigate } from 'react-router-dom';

export const Registration = () => {
  const isAuth = useSelector(selectIsAuth);
  const dispatch = useDispatch();
  const { 
    register,
    handleSubmit,
    formState:{errors,isValid} 
  } = useForm({
    defaultValues:{
      fullName:'Daulet Dauletov',
      email:'example@gmail.com',
      password:'123654789',
    },
    mode:"onChange"
  });
  const onSubmit = async(values) =>{
    const data = await dispatch(fetchRegister(values));
    if(!data.payload){
      return alert('Did not registred!');
    }
    if('token' in data.payload){
      window.localStorage.setItem('token',data.payload.token);
    }else{
      alert('Did not registred!');
    }
   
  };
  if(isAuth){
    return<Navigate to="/"/>;
  }
  return ( 
    <Paper classes={{ root: styles.root }}>
      <Typography classes={{ root: styles.title }} variant="h5">
        Creating account
      </Typography>
      <div className={styles.avatar}>
        <Avatar sx={{ width: 100, height: 100 }} />
      </div>
      <form onSubmit={handleSubmit(onSubmit)}>
      <TextField error={Boolean(errors.fullName?.message)}
        helperText={errors.fullName?.message}
        {...register('fullName', {required:'Write your full name'})}
         className={styles.field} label="Full name" fullWidth />

      <TextField error={Boolean(errors.email?.message)}
        helperText={errors.email?.message}
        type="email"
        {...register('email', {required:'Write your email'})}
        className={styles.field} label="E-Mail" fullWidth />

      <TextField className={styles.field} label="Password"
        error={Boolean(errors.password?.message)}
        helperText={errors.password?.message}
        {...register('password', {required:'Write password'})}
        fullWidth />
      <Button disabled={!isValid} type="submit" size="large" variant="contained" fullWidth>
        Register
      </Button>
      </form>
    </Paper>
  );
};

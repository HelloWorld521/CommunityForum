<template>
  <div class="register">
    <h2>Register</h2>
    <form @submit.prevent="onSubmit">
      <div>
        <label for="username">Username:</label>
        <input v-model="user.username" type="text" id="username" required/>
      </div>
      <div>
        <label for="email">Email:</label>
        <input v-model="user.email" type="email" id="email" required/>
      </div>
      <div>
        <label for="password">Password:</label>
        <input v-model="user.password" type="password" id="password" required/>
      </div>
      <button type="submit">Register</button>
    </form>
  </div>
</template>

<script>
import {computed, reactive} from 'vue';
import axios from 'axios';
import useVuelidate from '@vuelidate/core';
import {email, minLength, required} from '@vuelidate/validators';

export default {
  name: 'UserRegister', // 多词形式
  setup() {
    const user = reactive({
      username: '',
      email: '',
      password: ''
    });

    const rules = computed(() => {
      return {
        user: {
          username: {required},
          email: {required, email},
          password: {required, minLength: minLength(6)}
        }
      };
    });

    const v$ = useVuelidate(rules, {user});

    const onSubmit = () => {
      v$.value.$touch();
      if (!v$.value.$error) {
        axios.post('http://localhost:8081/api/v1/register', user)
            .then(() => {
              alert('Registration successful!');
              // Redirect to login or another page
              this.$router.push('/login'); // 使用 Vue Router 进行重定向
            })
            .catch(error => {
              console.error('There was an error!', error);
              // 根据后端返回的信息更新错误提示
              alert('Registration failed. ' + (error.response?.data.message || 'Please try again.'));
            });
      }
    };

    return {
      user,
      onSubmit,
      v$
    };
  }
};
</script>

<style scoped>
/* Add some basic styling */
.register {
  max-width: 400px;
  margin: 0 auto;
}

form div {
  margin-bottom: 15px;
}

input {
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
}

button {
  width: 100%;
  padding: 10px;
  background-color: #4CAF50;
  color: white;
  border: none;
  cursor: pointer;
}

button:hover {
  opacity: 0.8;
}
</style>
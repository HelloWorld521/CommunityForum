<template>
  <div class="login">
    <h2>Login</h2>
    <form @submit.prevent="onSubmit">
      <div>
        <label for="email">Email:</label>
        <input v-model="credentials.email" type="email" id="email" required />
      </div>
      <div>
        <label for="password">Password:</label> <!-- 修复for属性 -->
        <input v-model="credentials.password" type="password" id="password" required />
      </div>
      <button type="submit">Login</button>
    </form>
  </div>
</template>

<script>
import { reactive, computed } from 'vue';
import axios from 'axios';
import useVuelidate from '@vuelidate/core';
import { required, email } from '@vuelidate/validators';

export default {
  name: 'UserLogin', // 多词形式
  setup() {
    const credentials = reactive({
      email: '',
      password: ''
    });

    const rules = computed(() => {
      return {
        credentials: {
          email: { required, email },
          password: { required }
        }
      };
    });

    const v$ = useVuelidate(rules, { credentials });

    const onSubmit = () => {
      v$.value.$touch();
      if (!v$.value.$error) {
        axios.post('http://localhost:8081/api/v1/login', credentials)
          .then(response => {
            const token = response.data.token;
            localStorage.setItem('token', token);
            alert('Login successful!');
            // Redirect to a protected route or dashboard
          })
          .catch(error => {
            console.error('There was an error!', error);
            alert('Login failed. Please check your credentials and try again.');
          });
      }
    };

    return {
      credentials,
      onSubmit,
      v$
    };
  }
};
</script>

<style scoped>
/* Add some basic styling */
.login {
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
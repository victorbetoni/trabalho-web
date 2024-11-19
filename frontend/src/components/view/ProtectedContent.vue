<template>

</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { checkSession, logout } from '../../api/auth';

const router = useRouter();

checkSession((resp) => {
  if(resp.status != 200) {
    logout(async (_) => {
      router.push("/login");
      await router.go(0)
    })
  }
})

</script>
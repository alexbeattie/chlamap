<template>
  <div class="aba-centers-view">
    <div class="header">
      <h1>ABA Therapy Centers</h1>
      <p>Add and manage ABA therapy center information</p>
    </div>
    
    <div class="content">
      <ABACenterForm />
    </div>
  </div>
</template>

<script>
import ABACenterForm from '@/components/ABACenterForm.vue'

export default {
  name: 'ABACentersView',
  components: {
    ABACenterForm
  }
}
</script>

<style scoped>
.aba-centers-view {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.header {
  margin-bottom: 30px;
  text-align: center;
}

.header h1 {
  font-size: 2rem;
  color: #2c3e50;
  margin-bottom: 10px;
}

.header p {
  color: #666;
  font-size: 1.1rem;
}

.content {
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}
</style>
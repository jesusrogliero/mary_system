<template>
  <div class="home">
    <personas/>
  </div>
</template>

<script>
// @ is an alias to /src
import Personas from '@/components/Personas.vue'

export default {
  name: 'Persona',
  components: {
    Personas
  }
}
</script>

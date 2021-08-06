<template>
    <v-container>
       <v-simple-table
        fixed-header
        height="500px"
       >
            <template v-slot:default>
            <thead>
                <tr>
                    <th class="text-center">id</th>
                    <th class="text-center">Nombre</th>
                    <th class="text-center">Sexo</th>
                </tr>
            </thead>
            <tbody>
                <tr
                v-for="persona in personas"
                :key="persona.id"
                >
                <td>{{ persona.Id }}</td>
                <td>{{ persona.Name }}</td>
                <td>{{ persona.Sex }}</td>
                </tr>
            </tbody>
            </template>
        </v-simple-table>

   <v-row justify="center">
    <v-dialog
      v-model="dialog"
      persistent
      max-width="600px"
    >
      <v-card>
        <v-card-title>
          <span class="text-h5">User Profile</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col
                cols="6"
              >
                <v-text-field
                  label="Nombre*"
                  v-model="name"
                  solo
                ></v-text-field>
              </v-col>
              <v-col
                cols="6"
              >
                <v-text-field
                  label="Sexo*"
                  required
                  v-model="sex"
                  solo
                ></v-text-field>
              </v-col>
            </v-row>
          </v-container>
          <small>*indicates required field</small>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="blue darken-1"
            text
            @click="dialog = false"
          >
            Close
          </v-btn>

          <v-btn
            color="blue darken-1"
            text
            @click="createPersona"
          >
            Save
          </v-btn>

        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row> 
        <v-btn
        class="mt-7"
          color="primary"
          dark
          @click="openDialog"
        >
          Crear Persona
        </v-btn>

      <v-snackbar
        v-if="message != null"
        v-model="snackbar"
        :timeout="timeout"
      >
        {{ message }}
      </v-snackbar>
    </v-container>

</template>

<script>
export default {
    data: function() {
        return {
            personas: null,
            name: null,
            sex: null,
            message: null,
            dialog: null,
            snackbar: false,
            timeout: 2000
        }
    },

    watch: {
        dialog: function() {
            if(this.dialog == false)
              this.getPersona();
        }
    },

    mounted: function () {
        this.$nextTick(function () {
            this.getPersona();
        });
    },

    methods: {
        getPersona: function() {
            window.backend.getPersona().then(r => {
                this.personas = JSON.parse(r);
            });
        },

        createPersona: function() {
            window.backend.create(this.name, this.sex).then(r => {
                this.message = r;
            });

            this.dialog = null;
        },

        openDialog: function() {
            this.dialog = true;
        }
    }
}
</script>
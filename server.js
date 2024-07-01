const { ApolloServer } = require('@apollo/server');
const { ApolloGateway } = require('@apollo/gateway');

const gateway = new ApolloGateway({
  serviceList: [
    { name: 'inventory', url: 'http://localhost:8080/.well-known/apollo/server-health' },
    // Agrega más servicios si es necesario
  ],
});

const server = new ApolloServer({
  gateway,
  subscriptions: false, // Opcional, dependiendo de tu configuración de suscripciones
});

server.listen().then(({ url }) => {
  console.log(`🚀 Server ready at ${url}`);
});



/*const { ApolloServer } = require('@apollo/server');
const { ApolloGateway } = require('@apollo/gateway');
const { startStandaloneServer } = require('@apollo/server/standalone');
const fs = require('fs').promises;

async function startServer() {
  try {
    // Lee el contenido del archivo schema.graphql
    const schemaSdl = await fs.readFile('./graph/schema.graphql', 'utf-8');

    // Configura Apollo Gateway con el esquema cargado
    const gateway = new ApolloGateway({
      supergraphSdl: schemaSdl,
    });

    const server = new ApolloServer({ gateway, subscriptions: false });
  
    // Inicia el servidor independiente
    const { url } = await startStandaloneServer(server, {
      listen: { port: 4000 },
    });

    console.log(`🚀 Server ready at ${url}`);
  } catch (error) {
    console.error('Error starting server:', error);
  }
}

startServer();
*/
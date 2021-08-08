import SwaggerUI from 'swagger-ui-react';
import 'swagger-ui-react/swagger-ui.css';

export const ApiDesc = () => <SwaggerUI url={'/docs/api/public.yaml'} />;

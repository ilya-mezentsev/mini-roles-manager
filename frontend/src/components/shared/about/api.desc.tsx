import { useState } from 'react';
import ExpandMoreIcon from "@material-ui/icons/ExpandMore";
import {
    Accordion,
    AccordionDetails,
    AccordionSummary,
    Typography,
} from '@material-ui/core';

enum ApiExamples {
    CURL,
    JS,
    PYTHON,
}

export const ApiDesc = () => {
    const [expanded, setExpanded] = useState<ApiExamples | boolean>(ApiExamples.CURL);

    const handleChange = (panel: ApiExamples) => (_: any, isExpanded: boolean) => {
        setExpanded(isExpanded ? panel : false);
    };

    return (
        <>
            <h2>API</h2>
            <h3>
                Authorization:
            </h3>
            <p>
                Authorization is performed by API Token in HTTP headers. Header key is "X-RM-Auth-Token"
            </p>

            <h3>
                Request body:
            </h3>
            <ul>
                <li>roleId - Id of role that are going to perform operation under resource</li>
                <li>resourceId - Id of resource</li>
                <li>operation - Operation that are going to perform (create|read|update|delete)</li>
            </ul>

            <h3>
                Responses:
            </h3>
            <h4>
                Success:
            </h4>
            <p>HTTP 200 Ok</p>
            <pre>
                <code>
                    {"{"}"status": "ok", "data": {"{"}"effect": "permit"{"}"}{"}"}
                </code>
            </pre>
            <pre>
                <code>
                    {"{"}"status": "ok", "data": {"{"}"effect": "deny"{"}"}{"}"}
                </code>
            </pre>

            <h4>
                Bad request (Invalid request body):
            </h4>
            <p>400 Bad Request</p>
            <pre>
                <code>
                    Invalid JSON format
                </code>
            </pre>

            <h4>
                Unauthorized (No token in headers):
            </h4>
            <p>401 Unauthorized</p>
            <pre>
                <code>
                    {"{"}<br/>
                    &nbsp; "data": {"{"}<br/>
                    &nbsp; &nbsp; "code":"missed-token-in-headers",<br/>
                    &nbsp; &nbsp; "description":"No auth token in headers"<br/>
                    &nbsp; {"}"},<br/>
                    &nbsp; "status":"error"<br/>
                    {"}"}
                </code>
            </pre>

            <h4>
                Forbidden (Provided token does not exists):
            </h4>
            <p>403 Forbidden</p>
            <pre>
                <code>
                    {"{"}<br/>
                    &nbsp; "data":{"{"}<br/>
                    &nbsp; &nbsp; "code":"no-account-by-token",<br/>
                    &nbsp; &nbsp; "description":"Unable to find account by provided token"<br/>
                    &nbsp; {"}"},<br/>
                    &nbsp; "status":"error"<br/>
                    {"}"}
                </code>
            </pre>

            <h4>
                Server Error (Something horrible happened):
            </h4>
            <p>500 Internal Server Error</p>
            <pre>
                <code>
                    {"{"}<br/>
                    &nbsp; "data":{"{"}<br/>
                    &nbsp; &nbsp; "code":"unknown-error",<br/>
                    &nbsp; &nbsp; "description":"Unknown error happened"<br/>
                    &nbsp; {"}"},<br/>
                    &nbsp; "status":"error"<br/>
                    {"}"}
                </code>
            </pre>

            <h3>Examples:</h3>
            <Accordion expanded={expanded === ApiExamples.CURL} onChange={handleChange(ApiExamples.CURL)}>
                <AccordionSummary
                    expandIcon={<ExpandMoreIcon />}
                    aria-controls="panel1bh-content"
                    id="panel1bh-header"
                >
                    <Typography>CURL</Typography>
                </AccordionSummary>
                <AccordionDetails>
                    <Typography>
                        <pre>
                            <code>
                                curl -X POST localhost:8000/api/permissions \ <br/>
                                    -H "X-RM-Auth-Token: YOUR_API_TOKEN" \ <br/>
                                    -d '{"{"}"roleId": "role-1", "resourceId": "resource-1", "operation": "create"{"}"}'
                            </code>
                        </pre>
                    </Typography>
                </AccordionDetails>
            </Accordion>

            <Accordion expanded={expanded === ApiExamples.JS} onChange={handleChange(ApiExamples.JS)}>
                <AccordionSummary
                    expandIcon={<ExpandMoreIcon />}
                    aria-controls="panel1bh-content"
                    id="panel1bh-header"
                >
                    <Typography>JavaScript (Node.js)</Typography>
                </AccordionSummary>
                <AccordionDetails>
                    <Typography>
                        <pre>
                            <code>
                                const fetch = require('node-fetch'); <br/> <br/>

                                fetch(<br/>
                                    &nbsp; 'http://localhost:8000/api/permissions', <br/>
                                    &nbsp; {"{"} <br/>
                                    &nbsp; &nbsp; method: 'POST', <br/>
                                    &nbsp; &nbsp; headers: {"{"} <br/>
                                    &nbsp; &nbsp; &nbsp; 'X-RM-Auth-Token': 'YOUR_API_TOKEN', <br/>
                                    &nbsp; &nbsp; {"},"} <br/>
                                    &nbsp; &nbsp; body: JSON.stringify({"{"} <br/>
                                    &nbsp; &nbsp; &nbsp; roleId: 'role-1', <br/>
                                    &nbsp; &nbsp; &nbsp; resourceId: 'resource-1', <br/>
                                    &nbsp; &nbsp; &nbsp; operation: 'create' <br/>
                                    &nbsp; &nbsp; {"}"}, <br/>
                                );
                            </code>
                        </pre>
                    </Typography>
                </AccordionDetails>
            </Accordion>

            <Accordion expanded={expanded === ApiExamples.PYTHON} onChange={handleChange(ApiExamples.PYTHON)}>
                <AccordionSummary
                    expandIcon={<ExpandMoreIcon />}
                    aria-controls="panel1bh-content"
                    id="panel1bh-header"
                >
                    <Typography>Python</Typography>
                </AccordionSummary>
                <AccordionDetails>
                    <Typography>
                        <pre>
                            <code>
                                import requests
                                <br/> <br/>
                                requests.post( <br/>
                                    &nbsp; url='http://localhost:8000/api/permissions', <br/>
                                    &nbsp; headers={"{"}'X-RM-Auth-Token': 'YOUR_API_TOKEN'{"}"}, <br/>
                                    &nbsp; json={"{"} <br/>
                                    &nbsp; &nbsp; 'roleId': 'role-1', <br/>
                                    &nbsp; &nbsp; 'resourceId': 'resource-1', <br/>
                                    &nbsp; &nbsp; 'operation': 'create' <br/>
                                    &nbsp; {"}"}, <br/>
                                )
                            </code>
                        </pre>
                    </Typography>
                </AccordionDetails>
            </Accordion>
        </>
    );
}

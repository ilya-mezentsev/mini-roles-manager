import { Link } from '@material-ui/core';

export const Export = () => (
    <>
        <h2>To download your account data click button below:</h2>
        <Link
            download
            target="_blank"
            href={`${window.location.origin}/api/web-app/app-data/export`}
        >
            Download
        </Link>
    </>
);

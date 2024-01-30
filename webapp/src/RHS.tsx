import React from 'react';

import {getTest, postTest} from './client';

const RHS = () => {
    return (

        <div
            style={{
                display: 'flex',
                flexDirection: 'column',
                height: '100%',
                width: '100%',
            }}
        >
            <button
                onClick={() => {
                    getTest().
                        then((res) => {
                            console.log(res);
                        });
                }}
            >{'Get test'}</button>

            <button
                onClick={() => {
                    postTest().
                        then((res) => {
                            console.log(res);
                        });
                }}
            >{'Post test'}</button>
        </div>
    );
};

export {
    RHS,
};


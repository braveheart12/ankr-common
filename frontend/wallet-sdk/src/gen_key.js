import session25519 from 'session25519';
import {
    sha256
} from 'js-sha256';


const sum = (salt, callback) => {
    let key = Math.random().toString()
    session25519(salt, key, (err, keys) => {
        if (err !== null) {
            throw err;
        }

        callback(keys)
    });
}


export const gen_key = (salt, callback) => {
    sum(salt, (keys) => {
        console.log(sha256(keys.publicSignKey).substr(0, 40))
        const addr = sha256(keys.publicSignKey).substr(0, 40).toUpperCase();

        callback({
            public_key: keys.publicSignKeyBase64,
            private_key: keys.secretSignKeyBase64,
            addr: addr,
        })
    });
};
<?php

$json = file_get_contents('tmp.json');
echo base64UrlEncode($json);

// $json = file_get_contents('base64.txt');
// echo base64UrlDecode($json);

/**
 * Encodes string into "Base 64 Encoding with URL and Filename Safe Alphabet" (RFC 4648).
 *
 * > Note: Base 64 padding `=` may be at the end of the returned string.
 * > `=` is not transparent to URL encoding.
 *
 * @see https://tools.ietf.org/html/rfc4648#page-7
 * @param string $input the string to encode.
 * @return string encoded string.
 * @since 2.0.12
 */
function base64UrlEncode($input)
{
    return strtr(base64_encode($input), '+/', '-_');
}

/**
 * Decodes "Base 64 Encoding with URL and Filename Safe Alphabet" (RFC 4648).
 *
 * @see https://tools.ietf.org/html/rfc4648#page-7
 * @param string $input encoded string.
 * @return string decoded string.
 * @since 2.0.12
 */
function base64UrlDecode($input)
{
    return base64_decode(strtr($input, '-_', '+/'));
}
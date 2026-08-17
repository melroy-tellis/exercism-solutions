#include "complex_numbers.h"
#include <math.h>

complex_t c_add(complex_t a, complex_t b)
{
    complex_t sum = {c_real(a)+c_real(b), c_imag(a)+c_imag(b)};
    return sum;
}

complex_t c_sub(complex_t a, complex_t b)
{
    complex_t difference = {c_real(a)-c_real(b), c_imag(a)-c_imag(b)};
    return difference;
}

complex_t c_mul(complex_t a, complex_t b)
{
    complex_t product = {
        c_real(a)*c_real(b) - c_imag(a)*c_imag(b), 
        c_real(a)*c_imag(b) + c_imag(a)*c_real(b)
    };
    return product;
}

complex_t c_div(complex_t a, complex_t b)
{
    complex_t conjugate_b = c_conjugate(b);
    double z = c_abs(conjugate_b);
    double z2 = z*z;
    complex_t product_conjugate = c_mul(a, conjugate_b);
    complex_t quotient = {c_real(product_conjugate)/z2, c_imag(product_conjugate)/z2};
    return quotient;
}

double c_abs(complex_t x)
{
   return sqrt(c_real(x)*c_real(x) + c_imag(x)*c_imag(x));
}

complex_t c_conjugate(complex_t x)
{
    complex_t conjugate = {c_real(x), -c_imag(x)};
    return conjugate;
}

double c_real(complex_t x)
{
   return x.real;
}

double c_imag(complex_t x)
{
   return x.imag;
}

complex_t c_exp(complex_t x)
{
    double factor = exp(c_real(x));
    complex_t result = {factor*cos(c_imag(x)), factor*sin(c_imag(x))};
    return result;
    
}

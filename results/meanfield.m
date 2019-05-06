clear all;
close all;
clc;

data = csvread('meanfield_isingdata_CSV.csv', 1);

Temperature = data(:,1);
Magnetisation = data(:,2);
EnergyPerSite = data(:,3);
SpecificHeat = data(:,4);
MagneticSusceptibility = data(:,5);

figure;
plot(Temperature, Magnetisation, '.-')
xlabel('$\frac{T}{T_c}$', 'Interpreter', 'latex', 'FontSize', 16);
ylabel('$\mathcal{M}$', 'Interpreter', 'latex', 'FontSize', 16);
ax = gca;
ax.FontSize = 14; 

figure;
plot(Temperature, EnergyPerSite, '.-')
xlabel('$\frac{T}{T_c}$', 'Interpreter', 'latex', 'FontSize', 16);
ylabel('$\frac{E}{NJ}$', 'Interpreter', 'latex', 'FontSize', 16);
ax = gca;
ax.FontSize = 14; 

figure;
plot(Temperature, SpecificHeat, '.-')
xlabel('$\frac{T}{T_c}$', 'Interpreter', 'latex', 'FontSize', 16);
ylabel('$\frac{c}{k_B}$', 'Interpreter', 'latex', 'FontSize', 16);
ax = gca;
ax.FontSize = 14; 

figure;
plot(Temperature, MagneticSusceptibility, '.-')
xlabel('$\frac{T}{T_c}$', 'Interpreter', 'latex', 'FontSize', 16);
ylabel('$\chi$', 'Interpreter', 'latex', 'FontSize', 16);
ax = gca;
ax.FontSize = 14; 

% hold off;
% xlabel('$\frac{T}{T_c}$', 'Interpreter', 'latex', 'FontSize', 16);
% % ylabel('$\mathcal{M}$', 'Interpreter', 'latex', 'FontSize', 16);
% % ylabel('$\frac{E}{NJ}$', 'Interpreter', 'latex', 'FontSize', 16);
% % ylabel('$\frac{c}{k_B}$', 'Interpreter', 'latex', 'FontSize', 16);
% ylabel('$\chi$', 'Interpreter', 'latex', 'FontSize', 16);

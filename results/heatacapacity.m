clear all;
close all;
clc;

system = 0;
betaLowerLimit = 25;
betaUpperLimit = 100;
betaStep = 1;

betas = betaLowerLimit:betaStep:betaUpperLimit;

n0 = 2000
n = 50
m = (200000 - n0)/n

Ts = [];
HeatCapacities = [];
Chis = [];

for file = dir('awesome/*.csv')'    
    fname = ['./awesome/', file.name];
    data = csvread(fname, 1);

    Sweep = data(:,1);
    Temp = data(:,2);
    SubSystemID = data(:,3);
    Magnetisation = data(:,4);
    DimensionlessEnergy = data(:,5);
        
    magnetisations = Magnetisation(n0:n:(n0+n*m));
    energies = DimensionlessEnergy(n0:n:(n0+n*m));
    
    T = Temp(1)
    Ts = [Ts T]
        
    HeatCapacities = [HeatCapacities var(energies)/(T^2)];
    Chis = [Chis var(magnetisations)/T];
end

figure;
plot(Ts, HeatCapacities, '.')
xlabel('$T_0$', 'Interpreter', 'latex', 'FontSize', 16);
ylabel('Var($E$) / $T_0^2$', 'Interpreter', 'latex', 'FontSize', 16);

figure;
plot(Ts, Chis, '.')
xlabel('$T_0$', 'Interpreter', 'latex', 'FontSize', 16);
ylabel('Var($M$) / $T$', 'Interpreter', 'latex', 'FontSize', 16);

% ylim([-1, 0.2])

% legend_handle = legend('-DynamicLegend');
% set(legend_handle,'Interpreter','latex')
% legend('show');

 
% figure;
% plot(Sweep, DimensionlessEnergy)
% ylabel('$E/J$', 'Interpreter', 'latex', 'FontSize', 16);